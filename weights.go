package searcher

// 价值
type Weights struct {
	Weight float64
}

func NewWeights() *Weights {
	return &Weights{
		Weight: 0,
	}
}

func (t *Weights) AddValue(v float64) {
	t.Weight += v
}

func (t *Weights) Sum(w *Weights) *Weights {
	w0 := t.Clone()
	w0.AddValue(w.Weight)
	return w0
}

func (t *Weights) Clone() *Weights {
	return &Weights{
		Weight: t.Weight,
	}
}
