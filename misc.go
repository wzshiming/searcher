package searcher

import (
	"reflect"
	"regexp"
	"strings"
	"unicode"
)

func Addr(p interface{}) uintptr {
	return reflect.Indirect(reflect.ValueOf(p)).UnsafeAddr()
}

var word = regexp.MustCompile(`[^@#\$%\^&\*\\\/\(\)\[\]\{\}\s]+`)

func Grouping(str string) []string {
	d := GroupSpaces(str)
	r := []string{}

	for _, v := range d {
		r = append(r, GroupLanguages(v)...)
	}
	return r
}

// 空格和非法字符分组
func GroupSpaces(str string) []string {
	return word.FindAllString(str, -1)
}

// 不同语言之间分组
func GroupLanguages(str string) []string {
	d := []rune(str)
	switch len(d) {
	case 0:
		return []string{}
	case 1:
		return []string{str}
	default:
		b := d[0] < unicode.MaxASCII
		r := []string{string([]rune{d[0]})}
		for _, v := range d[1:] {
			if b == (v < unicode.MaxASCII) {
				r[len(r)-1] = string(append([]rune(r[len(r)-1]), v))
			} else {
				r = append(r, string([]rune{v}))
				b = !b
			}
		}
		return r
	}

	return nil
}

// 超简单的分词  列出分词所有可能性 单个词 按字拆分
func SimpleSegment(str string) (ret []string) {
	var w = []rune(str)
	for i := 0; i != len(w); i++ {
		for j := 0; j+i < len(w); j++ {
			v := string(w[j : j+i+1])
			ret = append(ret, v)
		}
	}
	return
}

// 超简单的分词  列出分词所有可能性 已经拆分好的词 组合所有可能性
func SimpleSegmentEndless(w []string) (ret []string) {
	for i := 0; i != len(w); i++ {
		for j := 0; j+i < len(w); j++ {
			v := strings.Join(w[j:j+i+1], "")
			ret = append(ret, v)
		}
	}
	return
}
