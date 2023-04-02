package gomamayo

import (
	"fmt"
	"strings"

	"github.com/goark/krconv"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

var vowel = map[string]string{"a": "ア", "i": "イ", "u": "ウ", "e": "エ", "o": "オ"}

type Analyzer struct {
	IsIgnored bool
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

// kagomeで解析する
func tokenize(input string) []tokenizer.Token {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
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

// New は Analyzer を作る
func New(isIgnored bool) *Analyzer {
	return &Analyzer{IsIgnored: isIgnored}
}

// Analyze は input がゴママヨか判定する
func (a Analyzer) Analyze(input string) (gomamayoResult GomamayoResult) {
	if a.IsIgnored {
		input, _ = applyIgnoreWordsRemoval(input)
	}

	tokens := tokenize(input)

	// for _, token := range tokens {
	// 	features := strings.Join(token.Features(), ",")
	// 	fmt.Printf("%s\t%v\n", token.Surface, features)
	// }

	vowelizedReading := []string{}

	for _, token := range tokens {
		reading, ok := token.Reading()
		if !ok {
			reading = token.Surface
		}
		if strings.Contains(reading, "ー") {
			reading = prolongedSoundMarkVowelize(reading)
		}
		vowelizedReading = append(vowelizedReading, reading)
	}

	for i := 0; i < len(tokens)-1; i++ {
		first := tokens[i]
		second := tokens[i+1]

		if first.POS()[0] != "名詞" && first.POS()[0] != "数詞" || second.Surface == first.Surface {
			continue
		}

		_, ok1 := first.Reading()
		_, ok2 := second.Reading()
		reading1 := vowelizedReading[i]
		reading2 := vowelizedReading[i+1]
		// 読みがある場合のみ
		// TODO: 読みがなくてもカナのみならそれを読みとするresultにも含めたい
		if ok1 && ok2 || reading1 != "" && reading2 != "" {
			minLen := min(len([]rune(reading1)), len([]rune(reading2)))
			for j := 1; j < minLen; j++ {
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
