package searcher

// 搜索类别
type Searcher struct {
	categorys map[string]*Searcher
	keywords  *Keywords
	parent    *Searcher
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

func (t *Searcher) Add(key string, p Point, v float64) {
	t.keywords.Add(key, p, v)
}

func (t *Searcher) Get(key string) *Values {
	ret := t.keywords.Get(key)
	for _, v := range t.categorys {
		ret = ret.UnionSet(v.Get(key))
	}
	return ret
}
