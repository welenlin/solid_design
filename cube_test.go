package shape

import (
	"math"
	"testing"
)

func TestCubeArea(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Cube Area: Cube with side 20 should return 2400",
			20,
			2400,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCube(tc.side)

			actual := c.GetArea()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}

func TestCubeCircumference(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Cube Circumference: Cube with side 20 should return 240",
			20,
			240,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCube(tc.side)

			actual := c.GetCircumference()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}
func TestCubeVolume(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Cube Volume: Cube with side 20 should return 8000",
			20,
			8000,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCube(tc.side)

			actual := c.GetVolume()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}
