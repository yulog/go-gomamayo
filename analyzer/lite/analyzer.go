package lite

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/yulog/go-gomamayo"
)

// 辞書を選択する
//
// https://github.com/ikawaha/kagome/blob/v2/cmd/tokenize/cmd.go
func selectDict(sysdict string) (*dict.Dict, int, error) {
	switch sysdict {
	case "ipa":
		return ipa.Dict(), ipa.Reading, nil
	case "uni", "uni2":
		return uni.Dict(), uni.LForm, nil
	}
	return nil, 0, fmt.Errorf("invalid dict name, %v", sysdict)
}

// New は Analyzer を作る
func New(sysdict string, isIgnored bool) (*gomamayo.Analyzer, error) {
	d, i, err := selectDict(sysdict)
	if err != nil {
		return nil, err
	}
	return &gomamayo.Analyzer{
		SysDict:      d,
		ReadingIndex: i,
		IsIgnored:    isIgnored,
	}, nil
}
