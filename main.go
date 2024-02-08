package main

import (
	"fmt"

	interfaceEx "github.com/MarkTBSS/go-structAndInterface/interface"
)

// ไม่ใช้ interface
type Rectangle struct {
	Width  float64
	Height float64
}

// ไม่ใช้ interface
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ไม่ใช้ interface
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Printf("Area of the rectangle: %.2f square units\n", area)
	fmt.Printf("Perimeter of the rectangle: %.2f units\n", perimeter)

	rectInterface := Rectangle{Width: 100, Height: 50}
	interfaceEx.PrintShapeDetails(rectInterface)
}
