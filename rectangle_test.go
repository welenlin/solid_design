package shape

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Rectangle Area: Rectangle with side 5 should return 25",
			5,
			25,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRectangle(tc.side)

			actual := r.GetArea()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}

func TestRectangleCircumference(t *testing.T) {
	tcs := []struct {
		name   string
		side   float64
		result float64
	}{
		{
			"Rectangle Circumference: Rectangle with side 5 should return 20",
			5,
			20,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := NewRectangle(tc.side)

			actual := r.GetCircumference()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}
