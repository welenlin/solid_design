package shape

type Shape interface {
	GetArea() float64
	GetCircumference() float64
}

type ThreeDimensionalShape interface {
	Shape
	GetVolume() float64
}
