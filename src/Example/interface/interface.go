package main

import (
	"fmt"
	"math"
)

// geometry
type geometry interface {
	area() float64      // 面积
	perimeter() float64 // 周长
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 实现方法 Rectangle面积
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// 实现方法 Rectangle周长
func (rect Rectangle) perimeter() float64 {
	return 2 * rect.width * rect.height
}

// 实现方法 Circle面积
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// 实现方法 Circle周长
func (circle Circle) perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area=", g.area())
	fmt.Println("Perimeter=", g.perimeter())
}
func main() {
	circle := Circle{5}
	fmt.Println("Area1=", circle.area())

	rectangle := Rectangle{10, 6}
	measure(circle)
	measure(rectangle)
}
