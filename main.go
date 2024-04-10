package main

import (
	"fmt"

	"github.com/MarkTBSS/go-structAndInterface/interface_example"
	"github.com/MarkTBSS/go-structAndInterface/no_interface"
)

func main() {
	rect := no_interface.Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	perimeter := rect.Perimeter()
	fmt.Printf("Area of the rectangle: %.2f square units\n", area)
	fmt.Printf("Perimeter of the rectangle: %.2f units\n", perimeter)

	rectInterface := interface_example.NewRectangle(100.00, 50.00)
	interface_example.PrintShapeDetails(rectInterface)
}
