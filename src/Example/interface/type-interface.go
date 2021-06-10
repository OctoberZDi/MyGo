package main

import "fmt"

type Square11 struct {
	side float32
}

type Circle11 struct {
	radius float32
}

func (s Square11) Area() float32 {
	panic("implement me")
}

func (s Circle11) Area() float32 {
	panic("implement me")
}

type Shaper11 interface {
	Area() float32
}

func main() {
	var areaIntf Shaper11
	sq1 := new(Square11)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square11); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle11); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
