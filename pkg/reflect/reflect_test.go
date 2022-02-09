package reflect

import (
	"fmt"
	"testing"
)

func TestReflect(t *testing.T) {
	type SS struct {
		A int     `json:"a"`
		B int8    `json:"b"`
		C int16   `json:"c"`
		D int32   `json:"d"`
		E int64   `json:"e"`
		F float64 `json:"f"`
		G bool    `json:"g"`
		H string  `json:"h"`
		I uint8   `json:"i"`
	}
	ss := &SS{}
	data := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
		"d": "4",
		"e": "5",
		"f": "6",
		"g": "true",
		"h": "8",
		"i": "9",
	}

	ss = Reflect(ss, data).(*SS)
	if !ss.G {
		t.Error("------")
	}

	fmt.Println("ss:", ss)
}
