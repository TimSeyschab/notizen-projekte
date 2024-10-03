package plotter

type Linear struct {
	Slope     float64
	Intercept float64
}

type Quadratic struct {
	A, B, C float64
}

type Cubic struct {
	A, B, C, D float64
}

func (l Linear) Plot(x int) float64 {
	return l.Slope*float64(x) + l.Intercept
}

func (q Quadratic) Plot(x int) float64 {
	xFloat := float64(x)
	return q.A*xFloat*xFloat + q.B*xFloat + q.C
}

func (c Cubic) Plot(x int) float64 {
	xFloat := float64(x)
	return c.A*xFloat*xFloat*xFloat + c.B*xFloat*xFloat + c.C*xFloat + c.D
}
