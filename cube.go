package shape

type Cube struct {
	side float64
}

const totalCubeSurface = 6

func NewCube(s float64) ThreeDimensionalShape {
	return Cube{
		side: s,
	}
}

func (c Cube) GetArea() float64 {
	return totalCubeSurface * c.side * c.side
}

func (c Cube) GetCircumference() float64 {
	return 2 * totalCubeSurface * c.side
}

func (c Cube) GetVolume() float64 {
	return c.side * c.side * c.side
}
