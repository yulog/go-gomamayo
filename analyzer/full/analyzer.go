package full

import (
	"fmt"

	ipaneologd "github.com/ikawaha/kagome-dict-ipa-neologd"
	uni3 "github.com/ikawaha/kagome-dict-uni3"
	"github.com/ikawaha/kagome-dict/dict"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/yulog/go-gomamayo"
)

// 辞書を選択する
//
// https://github.com/ikawaha/kagome/blob/v2/cmd/tokenize/cmd.go
func selectDict(sysdict string) (*dict.Dict, error) {
	switch sysdict {
	case "ipa":
		return ipa.Dict(), nil
	case "uni", "uni2":
		return uni.Dict(), nil
	case "uni3":
		return uni3.Dict(), nil
	case "neo", "neologd":
		return ipaneologd.Dict(), nil
	}
	return nil, fmt.Errorf("invalid dict name, %v", sysdict)
}

// New は Analyzer を作る
func New(sysdict string, isIgnored bool) (*gomamayo.Analyzer, error) {
	d, err := selectDict(sysdict)
	if err != nil {
		return nil, err
	}
	return &gomamayo.Analyzer{
		SysDict:   d,
		IsIgnored: isIgnored,
	}, nil
}
