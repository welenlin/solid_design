package printer

import (
	"fmt"

	"github.com/solid_design/shape"
)

type Printer interface {
	Print(s shape.Shape)
}

type printer struct {
}

// Inject Config to the parameters if exists for the initialization
func NewPrinter() Printer {
	return &printer{}
}

func (p *printer) Print(s shape.Shape) {
	fmt.Println("Shape Area: ", s.GetArea())
	fmt.Println("Shape Circumreference: ", s.GetCircumference())
}
