package shape

import (
	"math"
)

type Circle struct {
	diameter float64
}

func NewCircle(d float64) Shape {
	return Circle{
		diameter: d,
	}
}

func (c Circle) GetArea() float64 {
	var r = c.diameter / 2
	return math.Pi * r * r
}

func (c Circle) GetCircumference() float64 {
	return math.Pi * c.diameter
}
