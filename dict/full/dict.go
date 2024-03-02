package full

import (
	"fmt"

	ipaneologd "github.com/ikawaha/kagome-dict-ipa-neologd"
	uni3 "github.com/ikawaha/kagome-dict-uni3"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	"github.com/yulog/go-gomamayo/dict"
)

// 辞書を選択する
//
// https://github.com/ikawaha/kagome/blob/f4e6404b56ecf95b51836fbf2116406064007e7e/cmd/tokenize/cmd.go#L97
func SelectDict(sysdict string) (*dict.Dict, error) {
	switch sysdict {
	case "ipa":
		return &dict.Dict{SysDict: ipa.Dict(), ReadingIndex: ipa.Reading}, nil
	case "uni", "uni2":
		return &dict.Dict{SysDict: uni.Dict(), ReadingIndex: uni.LForm}, nil
	case "uni3":
		return &dict.Dict{SysDict: uni3.Dict(), ReadingIndex: uni3.LForm}, nil
	case "neo", "neologd":
		return &dict.Dict{SysDict: ipaneologd.Dict(), ReadingIndex: ipaneologd.Reading}, nil
	}
	return &dict.Dict{}, fmt.Errorf("invalid dict name, %v", sysdict)
}
