package shape

import (
	"math"
	"testing"
)

func TestSquareArea(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Square Area: Square with side 5 should return 25",
			5,
			25,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSquare(tc.side)

			actual := s.GetArea()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}

func TestSquareCircumference(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Square Circumference: Square with side 5 should return 20",
			5,
			20,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			s := NewSquare(tc.side)

			actual := s.GetCircumference()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}
