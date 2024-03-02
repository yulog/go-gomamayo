package gomamayo

import (
	"fmt"
	"strings"

	"github.com/goark/krconv"
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome/v2/filter"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var vowel = map[string]string{"a": "ア", "i": "イ", "u": "ウ", "e": "エ", "o": "オ"}

type Analyzer struct {
	SysDict *dict.Dict
	// Reading()はuni系だと使えないため、各辞書のconstから持って来て利用する
	// https://github.com/ikawaha/kagome-dict/blob/339b5b8724769ec9506e29fb5e771a9ec012784f/uni/tool/main.go#L39
	ReadingIndex int
	IsIgnored    bool
}

type GomamayoResult struct {
	IsGomamayo bool             `json:"isGomamayo"`
	Combo      int              `json:"combo"`
	Detail     []GomamayoDetail `json:"detail"`
}

type GomamayoDetail struct {
	Surface    string              `json:"surface"`
	Dimension  int                 `json:"dimension"`
	RawResult1 tokenizer.TokenData `json:"rawResult1"`
	RawResult2 tokenizer.TokenData `json:"rawResult2"`
}

// 辞書を選択する
//
// https://github.com/ikawaha/kagome/blob/v2/cmd/tokenize/cmd.go
// func selectDict(sysdict string) (*dict.Dict, error) {
// 	switch sysdict {
// 	case "ipa":
// 		return ipa.Dict(), nil
// 	case "neo", "neologd":
// 		return ipaneologd.Dict(), nil
// 	}
// 	return nil, fmt.Errorf("invalid dict name, %v", sysdict)
// }

// kagomeで解析する
func (a Analyzer) tokenize(input string) []tokenizer.Token {
	t, err := tokenizer.New(a.SysDict, tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	return t.Tokenize(input)
}

// 長音を変換する
//
// TODO: 直音化？https://github.com/goark/kkconv
func prolongedSoundMarkVowelize(reading string) (returnReading string) {
	prev := ""
	for _, readingRune := range reading {
		current := string(readingRune)
		if current != "ー" {
			returnReading += current
			prev = current
		} else if prev != "" {
			roman := []rune(krconv.Convert(prev))
			returnReading += vowel[string(roman[len(roman)-1])]
		}
	}
	return returnReading
}

// Deprecated: New は Analyzer を作る
// func New(sysdict string, isIgnored bool) (*Analyzer, error) {
// 	d, err := selectDict(sysdict)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Analyzer{
// 		SysDict:   d,
// 		IsIgnored: isIgnored,
// 	}, nil
// }

// Analyze は input がゴママヨか判定する
func (a Analyzer) Analyze(input string) (gomamayoResult GomamayoResult) {
	if a.IsIgnored {
		input, _ = applyIgnoreWordsRemoval(input)
	}

	tokens := a.tokenize(input)

	posFilterAllow := filter.NewPOSFilter([]filter.POS{
		{"名詞"},
	}...)
	posFilterNotAllow := filter.NewPOSFilter([]filter.POS{
		{filter.Any, "数詞"},
	}...)

	// for _, token := range tokens {
	// 	features := strings.Join(token.Features(), ",")
	// 	fmt.Printf("%s\t%v\n", token.Surface, features)
	// }

	vowelizedReading := []string{}

	for _, token := range tokens {
		reading, ok := token.FeatureAt(a.ReadingIndex) // 読みを取得
		if !ok {
			reading = token.Surface // 読みがない場合、Surfaceを読みとする
		}
		if strings.Contains(reading, "ー") {
			reading = prolongedSoundMarkVowelize(reading)
		}
		vowelizedReading = append(vowelizedReading, reading)
	}

	for i := 0; i < len(tokens)-1; i++ {
		first := tokens[i]
		second := tokens[i+1]

		// ThinaticSystem/gomamayo.jsとna2na-p/gomamayo-denoで条件が違う…？
		if !posFilterAllow.Match(first.POS()) || posFilterNotAllow.Match(first.POS()) || second.Surface == first.Surface {
			continue
		}

		_, ok1 := first.FeatureAt(a.ReadingIndex)  // 読みを取得
		_, ok2 := second.FeatureAt(a.ReadingIndex) // 読みを取得
		reading1 := vowelizedReading[i]
		reading2 := vowelizedReading[i+1]
		// 読みがある場合のみ
		// TODO: Surfaceを読みとしたならresultにも含めたい
		if ok1 && ok2 || reading1 != "" && reading2 != "" {
			minLen := min(len([]rune(reading1)), len([]rune(reading2)))
			for j := 1; j <= minLen; j++ {
				fragment1 := string([]rune(reading1)[len([]rune(reading1))-j:])
				fragment2 := string([]rune(reading2)[0:j])
				if fragment1 == fragment2 {
					gomamayoResult.IsGomamayo = true
					gomamayoResult.Detail = append(gomamayoResult.Detail, GomamayoDetail{
						Surface:    fmt.Sprintf("%s|%s", first.Surface, second.Surface),
						Dimension:  j,
						RawResult1: tokenizer.NewTokenData(first),
						RawResult2: tokenizer.NewTokenData(second),
					})
					gomamayoResult.Combo++
				}
			}
		}
	}

	return gomamayoResult
}
