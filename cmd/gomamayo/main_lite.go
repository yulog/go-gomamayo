//go:build lite

package main

import (
	"github.com/yulog/go-gomamayo/dict"
	"github.com/yulog/go-gomamayo/dict/lite"
)

func selectDict(sysdict string) (*dict.Dict, error) {
	d, err := lite.SelectDict(sysdict)
	if err != nil {
		return nil, err
	}
	return d, nil
}
