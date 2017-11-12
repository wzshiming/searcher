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

func (t *Values) Len() int {
	return len(t.maps)
}

func (t *Values) Add(p Point, v float64) {
	w := t.Get(p)
	w.AddValue(v)
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
func (t *Values) UnionSet(vas ...*Values) *Values {
	vs := t.Clone()
	for _, va := range vas {
		for k, v := range va.maps {
			m, ok := t.maps[k]
			if ok {
				vs.maps[k] = m.Sum(v)
			} else {
				vs.maps[k] = v.Clone()
			}
		}
	}
	return vs
}

// 交集
func (t *Values) IntersectionSet(vas ...*Values) *Values {
	vs := NewValues()
	for _, va := range vas {
		for k, v := range va.maps {
			m, ok := t.maps[k]
			if ok {
				vs.maps[k] = m.Sum(v)
			}
		}
	}
	return vs
}

func (t *Values) Data() (ps Points) {
	for k, _ := range t.maps {
		ps = append(ps, k)
	}
	return
}

func (t *Values) Sort(f func(Point, *Weights) float64) (ps Points) {
	ps = t.Data()
	ps.Sort(func(a Point, b Point) bool {
		aw := f(a, t.Get(a))
		bw := f(b, t.Get(b))
		if aw == bw {
			return Addr(a) < Addr(b)
		}
		return aw < bw
	})
	return
}

func (t *Values) Clone() *Values {
	vs := NewValues()
	for k, v := range t.maps {
		vs.maps[k] = v.Clone()
	}
	return vs
}
