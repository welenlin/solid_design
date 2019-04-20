package shape

const rectangleTotalSide = 4

type Rectangle struct {
	side float64
}

func NewRectangle(s float64) Shape {
	return Rectangle{
		side: s,
	}
}
func (r Rectangle) GetArea() float64 {
	return r.side * r.side
}

func (r Rectangle) GetCircumference() float64 {
	return rectangleTotalSide * r.side
}
