package searcher

import "sort"

type Sorter struct {
	less func(a Point, b Point) bool
	data Points
}

func NewSorter(data Points, less func(a Point, b Point) bool) *Sorter {
	return &Sorter{
		data: data,
		less: less,
	}
}

func (t *Sorter) Len() int {
	return len(t.data)
}

func (t *Sorter) Less(i, j int) bool {
	return t.less(t.data[i], t.data[j])
}

func (t *Sorter) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}

func (t *Sorter) Sort() {
	sort.Sort(t)
}
