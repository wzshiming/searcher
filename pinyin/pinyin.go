package pinyin

import (
	pinyin "github.com/mozillazg/go-pinyin"
	"github.com/wzshiming/searcher"
)

const DriverName = "pinyin"

func init() {
	searcher.RegisterDriver(DriverName, func(t *searcher.Searcher, key string, p searcher.Point, value float64) {
		w := ToPinYin(key)

		// 全拼
		d := searcher.SimpleSegmentEndless(w)
		for _, v := range d {
			t.AddWord(v, p, value)
		}

		// 首字母
		f := ""
		for _, v := range w {
			f += v[:1]
		}
		if f != "" {
			t.Add(f, p, value)
		}
	})
}

func ToPinYin(han string) []string {
	a := pinyin.NewArgs()
	var d = pinyin.Pinyin(han, a)
	var r []string
	for _, v := range d {
		if len(v) < 1 {
			continue
		}
		r = append(r, v[0])
	}
	return r
}
