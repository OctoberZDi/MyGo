package main

import (
	"fmt"
	"math"
)

// geometry 几何
type geometry interface {
	area() float64      // 面积
	perimeter() float64 // 周长
}

type Rectangle1 struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 实现方法 Rectangle1面积
func (rec Rectangle1) area() float64 {
	return rec.width * rec.height
}

// 实现方法 Rectangle1周长
func (rec Rectangle1) perimeter() float64 {
	return 2 * rec.width * rec.height
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

	Rectangle1 := Rectangle1{10, 6}
	measure(circle)
	measure(Rectangle1)
}
