package searcher

var driver = map[string]func(t *Searcher, key string, p Point, value float64){}

func RegisterDriver(name string, f func(t *Searcher, key string, p Point, value float64)) {
	driver[name] = f
}
