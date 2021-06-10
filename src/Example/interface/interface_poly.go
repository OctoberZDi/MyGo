package main

import "fmt"

// Shaper 定义一个接口
type Shaper interface {
	Area() float32
}

// Square 定义一个结构体
type Square struct {
	side float32
}

// Area 实现Area接口，求取正方形的面积
func (s Square) Area() float32 {
	return s.side * s.side
}

// Rectangle 定义一个结构体
type Rectangle struct {
	height, width float32
}

// Area 实现Area接口，求取长方形的面积
func (r Rectangle) Area() float32 {
	return r.height * r.width
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
