package gotopinyin

import (
	"regexp"
	"strconv"
	"strings"
)

func toHex(x int64) string {
	return strconv.FormatInt(x, 16)
}

func replaceWith(r *regexp.Regexp, s, with string) string {
	return string(r.ReplaceAll([]byte(s), []byte(with)))
}

func Convert(s, join string) string {
	pyStr := ""
	runes := []rune(s)
	digitalReg := regexp.MustCompile("\\d")

	for i, runesLen := 0, len(runes); i < runesLen; i++ {
		x := int64(runes[i])
		var chr string
		if x > 255 {
			sHex := strings.ToUpper(toHex(x))
			if pyList := DICT[sHex]; pyList != "" {
				chr = replaceWith(digitalReg, pyList, "")
			} else {
				chr = string(x)
			}
			if i == 0 {
				pyStr += chr
				if runesLen > (i + 1) {
					if runes[i+1] <= 255 {
						if runes[i+1] != 32 {
							pyStr += join
						}
					}
				}
			} else {
				if i != 0 && runes[i-1] != 32 {
					pyStr += join
				}
				pyStr += chr
				if runesLen > (i + 1) {
					if runes[i+1] <= 255 {
						if runes[i+1] != 32 {
							pyStr += join
						}
					}
				}
			}
		} else {
			pyStr += string(x)
		}
	}
	return pyStr
}
