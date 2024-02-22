package interfaceEx

import (
	"fmt"
)

// ใช้ interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ใช้ struct Rectangle และ implement interface Shape
type Rectangle struct {
	Width  float64
	Height float64
}

// Constructor
func NewRectangle(width, height float64) Shape {
	return &Rectangle{
		Width:  width,
		Height: height,
	}
}

// Method สำหรับคำนวณพื้นที่ของ Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method สำหรับคำนวณเส้นรอบรูปของ Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// ใช้ฟังก์ชัน printShapeDetails เพื่อพิมพ์รายละเอียดของรูปทรงที่ได้รับ interface Shape
func PrintShapeDetails(s Shape) {
	area := s.Area()
	perimeter := s.Perimeter()

	fmt.Printf("Area of the shape: %.2f square units\n", area)
	fmt.Printf("Perimeter of the shape: %.2f units\n", perimeter)
}
