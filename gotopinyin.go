package	gotopinyin

import (
	"strconv"
	"regexp"
	"io/ioutil"
	"strings"
	"path"
	"runtime"
)

var DICT = map[string][]string{}
var dictLoaded bool

func initDict() {
	if dictLoaded { return }
	_, currentFile, _, _ := runtime.Caller(1)
	file := path.Join(path.Dir(currentFile), "pinyin.txt")
	b, err := ioutil.ReadFile(file)
	if err != nil { panic(err) }
	content := string(b)
	ary := make([]string, 1000)
	ary = strings.Split(content, "\n")
	reg := regexp.MustCompile("\\s|\\(|\\)")
	for _, str := range ary {
		tmp := strings.Split(str, " ")
		key, value := tmp[0], tmp[1]
		valueSanitized := replaceWith(reg, value, "")
		valueAry := strings.Split(valueSanitized, ",")
		collect := make([]string, 0)
		for i:=0; i< len(valueAry); i++ {
			collect = append(collect, valueAry[i])
		}
		DICT[key] = collect
	}
	dictLoaded = true
}

func toHex(x int64) string {
	return strconv.FormatInt(x, 16)
}

func replaceWith(r *regexp.Regexp, s ,with string) string {
	return string(r.ReplaceAll([]byte(s), []byte(with)))
}

func Convert(s, join string) string {
	initDict()
	pyStr := ""
	runes := []rune(s)
	digitalReg := regexp.MustCompile("\\d")

	for i, runesLen := 0, len(runes); i < runesLen; i++ {
		x := int64(runes[i])
		var chr string
		if x > 255 {
			sHex := strings.ToUpper(toHex(x))
			if pyAry := DICT[sHex]; pyAry != nil {
				chr = replaceWith(digitalReg, pyAry[0], "")
			} else {
				chr = string(x)
			}
			if i == 0 {
				pyStr += chr
				if runesLen > (i+1) {
					if runes[i + 1] <= 255 {
						if runes[i + 1] != 32 {
							pyStr += join
						}
					}
				}
			} else {
				if i != 0 && runes[i - 1] != 32 {
					pyStr += join
				}
				pyStr += chr
				if runesLen > (i+1) {
					if runes[i + 1] <= 255 {
						if runes[i + 1] != 32 {
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
