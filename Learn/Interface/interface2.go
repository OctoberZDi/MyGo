package main

import (
	"fmt"
	"math"
)

// Shape 形状
type Shape interface {
	area() float64
}

// Circle 圆形
type Circle struct {
	x, y, radius float64
}

// Rectangle 长方形
type Rectangle struct {
	width, height float64
}

// define a method for Circle(implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// define a method for rectangle(implementation of Shape.area())
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 5}
	fmt.Printf("Circle area: %10f \n", getArea(circle))
	fmt.Printf("Rectangle area: %10f \n", getArea(rectangle))
}
