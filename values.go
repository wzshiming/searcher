package searcher

// 某关键字下索引的数据权值
type Values struct {
	maps map[Point]*Weights
}

func NewValues() *Values {
	return &Values{
		maps: map[Point]*Weights{},
	}
}

func (t *Values) Add(p Point, v float64) {
	t.Get(p).AddValue(v)
}

func (t *Values) Del(p Point) {
	delete(t.maps, p)
}

func (t *Values) Get(p Point) *Weights {
	w, ok := t.maps[p]
	if !ok {
		w = NewWeights()
		t.maps[p] = w
	}
	return w
}

// 并集
func (t *Values) UnionSet(va *Values) *Values {
	vs := t.Clone()
	for k, v := range va.maps {
		m, ok := vs.maps[k]
		if ok {
			vs.maps[k] = m.Sum(va.maps[k])
		}
		vs.maps[k] = v
	}
	return vs
}

// 交集
func (t *Values) IntersectionSet(va *Values) *Values {
	vs := t.Clone()
	for k, _ := range va.maps {
		m, ok := vs.maps[k]
		if ok {
			vs.maps[k] = m.Sum(va.maps[k])
		}
		vs.Del(k)
	}
	return vs
}

func (t *Values) Sort(f func(Point, *Weights) float64) (ps Points) {
	for k, _ := range t.maps {
		ps = append(ps, k)
	}
	ps.Sort(func(a Point, b Point) bool {
		return f(a, t.Get(a)) < f(b, t.Get(b))
	})
	return
}

func (t *Values) Clone() *Values {
	vs := NewValues()
	for k, v := range t.maps {
		vs.maps[k] = v
	}
	return vs
}
