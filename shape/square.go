package shape

const squareTotalSide = 4

type Square struct {
	side float64
}

func NewSquare(s float64) Shape {
	return Square{
		side: s,
	}
}
func (r Square) GetArea() float64 {
	return r.side * r.side
}

func (r Square) GetCircumference() float64 {
	return squareTotalSide * r.side
}
