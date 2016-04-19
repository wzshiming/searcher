package searcher

// 被索引的数据
type Point interface{}

type Points []Point

func (t *Points) Sort(less func(Point, Point) bool) {
	NewSorter(*t, less).Sort()
}
