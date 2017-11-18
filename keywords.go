package searcher

import (
	"strings"
)

// 固定类别下全部关键字对应的索引
type Keywords struct {
	keys map[string]*Values
}

func NewKeywords() *Keywords {
	return &Keywords{
		keys: map[string]*Values{},
	}
}

func (t *Keywords) Add(key string, p Point, value float64) {
	ks := Grouping(key)
	keys := []string{}
	for _, v := range ks {
		keys = append(keys, SimpleSegment(v)...)
	}
	kl := float64(len([]rune(key)))
	for _, v := range keys {
		vl := float64(len([]rune(v)))
		v0 := value * vl / kl
		t.AddWord(v, p, v0)
	}
}

func (t *Keywords) AddWord(key string, p Point, value float64) {
	key = strings.ToLower(key)
	v, ok := t.keys[key]
	if !ok {
		v = NewValues()
		t.keys[key] = v
	}
	v.Add(p, value)
}

//func (t *Keywords) Get(key string) *Values {
//	ks := Grouping(key)
//	vs := NewValues()
//	for _, v := range ks {
//		vs = vs.UnionSet(t.GetWord(v))
//	}
//	return vs
//}

func (t *Keywords) GetWord(key string) *Values {
	key = strings.ToLower(key)
	v, ok := t.keys[key]
	if !ok {
		return NewValues()
	}
	return v
}
