package structsandinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circ Circle) Area() float64 {
	return math.Pi * circ.Radius * circ.Radius
}

func (tria Triangle) Area() float64 {
	return (tria.Height * tria.Base) / 2
}
