package gotopinyin

import (
	"testing"
)

var fixturesWithSpace = map[string]string {
	"": "",
	" ": " ",
	"a": "a",
	" a": " a",
	"a ": "a ",
	" a ": " a ",
	"abc": "abc",
	"a bc": "a bc",
	"汉": "han",
	"汉 ": "han ",
	" 汉": " han",
	"汉 字": "han zi",
	" 汉 字": " han zi",
	" 汉 字 ": " han zi ",
	"HAN汉ZI字": "HAN han ZI zi",
}

var fixturesWithDash = map[string]string {
	"": "",
	" ": " ",
	"a": "a",
	" a": " a",
	"a ": "a ",
	" a ": " a ",
	"abc": "abc",
	"a bc": "a bc",
	"汉": "han",
	"汉 ": "han ",
	" 汉": " han",
	"汉 字": "han zi",
	" 汉 字": " han zi",
	" 汉 字 ": " han zi ",
	"HAN汉ZI字": "HAN-han-ZI-zi",
	"HAN汉字ZI": "HAN-han-zi-ZI",
}

func TestJoinWithSpace (t *testing.T) {
	for k, v := range fixturesWithSpace {
		r := Convert(k, " ")
		if r != v {
			t.Errorf("(%s)(%s)", r, v)
		}
	}
}

func TestJoinWithDash (t *testing.T) {
	for k, v := range fixturesWithDash {
		r := Convert(k, "-")
		if r != v {
			t.Errorf("(%s)(%s)", r, v)
		}
	}
}
