package searcher

import (
	"sync"
)

// 搜索类别
type Searcher struct {
	categorys map[string]*Searcher
	keywords  *Keywords
	parent    *Searcher
	mut       sync.RWMutex
}

func newSearcher(parent *Searcher) *Searcher {
	return &Searcher{
		categorys: map[string]*Searcher{},
		parent:    parent,
		keywords:  NewKeywords(),
	}
}

func NewSearcher() *Searcher {
	return newSearcher(nil)
}

func (t *Searcher) Category(category string) *Searcher {
	c, ok := t.categorys[category]
	if !ok {
		c = newSearcher(t)
		t.categorys[category] = c
	}
	return c
}

func (t *Searcher) AddBy(drivername string, key string, p Point, value float64) {
	f := driver[drivername]
	if f == nil {
		return
	}
	f(t, key, p, value)
}

func (t *Searcher) Add(key string, p Point, value float64) {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.keywords.Add(key, p, value)
}

func (t *Searcher) AddWord(key string, p Point, value float64) {
	t.mut.Lock()
	defer t.mut.Unlock()
	t.keywords.AddWord(key, p, value)
}

func (t *Searcher) Get(key string) *Values {
	t.mut.RLock()
	defer t.mut.RUnlock()
	ks := Grouping(key)
	rets0 := []*Values{}
	for _, k := range ks {
		rets := []*Values{}
		d := t.keywords.GetWord(k)
		if d.Len() > 0 {
			rets = append(rets, d)
		}
		for _, v := range t.categorys {
			d := v.Get(k)
			if d.Len() > 0 {
				rets = append(rets, d)
			}
		}
		switch len(rets) {
		case 0:
			return NewValues()
		case 1:
			rets0 = append(rets0, rets[0].Clone())
		default:
			rets0 = append(rets0, rets[0].UnionSet(rets[1:]...))
		}
	}

	switch len(rets0) {
	case 0:
		return NewValues()
	case 1:
		return rets0[0]
	default:
		return rets0[0].IntersectionSet(rets0[1:]...)
	}
}
