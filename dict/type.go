package dict

import (
	"github.com/ikawaha/kagome-dict/dict"
)

type Dict struct {
	SysDict *dict.Dict
	// Reading()はuni系だと使えないため、各辞書のconstから持って来て利用する
	// https://github.com/ikawaha/kagome-dict/blob/339b5b8724769ec9506e29fb5e771a9ec012784f/uni/tool/main.go#L39
	ReadingIndex int
}
