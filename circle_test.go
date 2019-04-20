package shape

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	tcs := []struct {
		name     string
		diameter float64
		result   float64
	}{
		{
			"Circle Area: Circle with diameter 14 should return 154",
			14,
			154,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCircle(tc.diameter)

			actual := c.GetArea()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}

func TestCircleCircumference(t *testing.T) {
	tcs := []struct {
		name     string
		diameter float64
		result   float64
	}{
		{
			"Circle Circumference: Circle with diameter 14 should return 4",
			14,
			44,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			c := NewCircle(tc.diameter)

			actual := c.GetCircumference()

			if math.RoundToEven(tc.result) != math.RoundToEven(actual) {
				t.Errorf("Expected %f, Actual Result %f", math.RoundToEven(tc.result), math.RoundToEven(actual))
			}
		})
	}
}
