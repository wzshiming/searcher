// +build go1.9

package searcher

import "sort"

func (t *Points) Sort(less func(Point, Point) bool) {
	sort.Slice(*t, func(i, j int) bool {
		return less((*t)[i], (*t)[j])
	})
}
