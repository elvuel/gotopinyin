package gotopinyin

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	digitalReg = regexp.MustCompile("\\d")
)

func toHex(x int64) string {
	return strconv.FormatInt(x, 16)
}

func replaceWith(s, with string) string {
	return string(digitalReg.ReplaceAll([]byte(s), []byte(with)))
}

func Convert(s, join string) string {
	if !utf8.ValidString(s) {
		return ""
	}

	var buf bytes.Buffer
	runes := []rune(s)

	var x int64
	var chr, sHex, pyList string
	for i, runesLen := 0, len(runes); i < runesLen; i++ {
		x = int64(runes[i])
		if x > 255 {
			sHex = strings.ToUpper(toHex(x))
			if pyList = DICT[sHex]; pyList != "" {
				chr = replaceWith(pyList, "")
			} else {
				chr = string(x)
			}
			if i == 0 {
				buf.WriteString(chr)
				if runesLen > (i + 1) {
					if runes[i+1] <= 255 {
						if runes[i+1] != 32 {
							buf.WriteString(join)
						}
					}
				}
			} else {
				if i != 0 && runes[i-1] != 32 {
					buf.WriteString(join)
				}
				buf.WriteString(chr)
				if runesLen > (i + 1) {
					if runes[i+1] <= 255 {
						if runes[i+1] != 32 {
							buf.WriteString(join)
						}
					}
				}
			}
		} else {
			buf.WriteString(string(x))
		}
	}
	return buf.String()
}
