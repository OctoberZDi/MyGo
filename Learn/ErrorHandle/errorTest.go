package main

import (
	"errors"
	"fmt"
	"math"
)

type error interface {
	Error() string
}

//错误处理
func main() {
	fmt.Println(sqrt(-1))
	fmt.Println(sqrt(64))
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math:square root of negative number...")
	}

	return math.Sqrt(f), nil
}
