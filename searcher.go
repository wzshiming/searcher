package searcher

// 搜索类别
type Searcher struct {
	categorys map[string]*Keywords
}

func NewSearcher() *Searcher {
	return &Searcher{
		categorys: map[string]*Keywords{},
	}
}

func (t *Searcher) Category(category string) *Keywords {
	c, ok := t.categorys[category]
	if !ok {
		c = NewKeywords()
		t.categorys[category] = c
	}
	return c
}

func (t *Searcher) Add(key string, p Point, v float64) {
	t.Category("").Add(key, p, v)
}

func (t *Searcher) Get(key string) *Values {
	return t.Category("").Get(key)
}
