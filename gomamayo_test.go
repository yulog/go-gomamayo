package gomamayo

import (
	"testing"
)

func FuzzAnalyze(f *testing.F) {
	testcases := []string{"ごまマヨネーズ", " ", "メモリリーク", "メモリ、リーク。", "独自実装", "ダイレクト投稿", "部分分数分解"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		New(false).Analyze(s)
	})
}
