package surface

import "math"

type surfaceGraphFunc interface {
	f(x, y float64) float64
}

type defaultGraph struct{}

func (d *defaultGraph) f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

type eggBox struct{}

func (eb *eggBox) f(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 10
}

type saddle struct{}

func (s *saddle) f(x, y float64) float64 {
	return (x*x/4 - y*y/9) / 100
}

type mogul struct{}

func (m *mogul) f(x, y float64) float64 {
	return (math.Sin(x) + math.Cos(y)) / 10
}

func chooseSurfaceGraphFunc(sgn string) surfaceGraphFunc {
	switch sgn {
	case "eggbox":
		return &eggBox{}
	case "saddle":
		return &saddle{}
	case "mogul":
		return &mogul{}
	default:
		return &defaultGraph{}
	}
}
