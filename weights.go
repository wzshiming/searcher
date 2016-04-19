package searcher

import (
	"time"
)

// 价值
type Weights struct {
	Weight     float64
	CreateTime time.Time
	UpdateTime time.Time
}

func NewWeights() *Weights {
	return &Weights{
		Weight:     0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

func (t *Weights) AddValue(v float64) {
	t.Weight += v
	t.UpdateTime = time.Now()
}

func (t *Weights) Sum(w *Weights) *Weights {
	r := Weights{
		Weight: t.Weight + w.Weight,
	}
	if t.CreateTime.Before(w.CreateTime) {
		r.CreateTime = t.CreateTime
	} else {
		r.CreateTime = w.CreateTime
	}
	if t.UpdateTime.Before(w.UpdateTime) {
		r.UpdateTime = w.UpdateTime
	} else {
		r.UpdateTime = t.UpdateTime
	}
	return &r
}
