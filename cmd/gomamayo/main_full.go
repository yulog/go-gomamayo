//go:build !lite

package main

import (
	"github.com/yulog/go-gomamayo/dict"
	"github.com/yulog/go-gomamayo/dict/full"
)

func selectDict(sysdict string) (*dict.Dict, error) {
	d, err := full.SelectDict(sysdict)
	if err != nil {
		return nil, err
	}
	return d, nil
}
