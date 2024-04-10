package interface_example

import (
	"fmt"
)

type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Constructor
func NewRectangle(width, height float64) Shaper {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// ใช้ฟังก์ชัน printShapeDetails เพื่อพิมพ์รายละเอียดของรูปทรงที่ได้รับ interface Shape
func PrintShapeDetails(s Shaper) {
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Printf("Area of the shape: %.2f square units\n", area)
	fmt.Printf("Perimeter of the shape: %.2f units\n", perimeter)
}
