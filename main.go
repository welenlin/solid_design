package main

import (
	prnt "github.com/solid_design/printer"
	"github.com/solid_design/shape"
)

func main() {
	var circleDiameter float64 = 7
	c := shape.NewCircle(circleDiameter)
	p := prnt.NewPrinter()

	p.Print(c)
}
