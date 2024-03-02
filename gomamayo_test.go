package gomamayo

import (
	"testing"

	"github.com/yulog/go-gomamayo/dict/full"
)

func FuzzAnalyzeIpa(f *testing.F) {
	testcases := []string{"ごまマヨネーズ", " ", "メモリリーク", "メモリ、リーク。", "独自実装", "ダイレクト投稿", "部分分数分解", "福山雅治", "長期金利", "オレンジレンジ", "太鼓公募募集終了", "多項高次ゴママヨ"}
	d, _ := full.SelectDict("ipa")
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		New(d, false).Analyze(s)
	})
}

func FuzzAnalyzeUni(f *testing.F) {
	testcases := []string{"ごまマヨネーズ", " ", "メモリリーク", "メモリ、リーク。", "独自実装", "ダイレクト投稿", "部分分数分解", "福山雅治", "長期金利", "オレンジレンジ", "太鼓公募募集終了", "多項高次ゴママヨ"}
	d, _ := full.SelectDict("uni")
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		New(d, false).Analyze(s)
	})
}
