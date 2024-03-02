package lite

import (
	"fmt"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome-dict/uni"
	dictlib "github.com/yulog/go-gomamayo/dict"
)

// 辞書を選択する
//
// https://github.com/ikawaha/kagome/blob/f4e6404b56ecf95b51836fbf2116406064007e7e/cmd/tokenize/cmd.go#L97
func SelectDict(sysdict string) (*dictlib.Dict, error) {
	switch sysdict {
	case "ipa":
		return &dictlib.Dict{SysDict: ipa.Dict(), ReadingIndex: ipa.Reading}, nil
	case "uni", "uni2":
		return &dictlib.Dict{SysDict: uni.Dict(), ReadingIndex: uni.LForm}, nil
	}
	return &dictlib.Dict{}, fmt.Errorf("invalid dict name, %v", sysdict)
}
