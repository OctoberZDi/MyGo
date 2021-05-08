package main

import "fmt"

func main() {

	fmt.Printf("15的阶乘： %d", factorial(15))
}

func factorial(n uint64) (result uint64) {
	if n > 0 {
		return n * factorial(n-1)
	}

	return 1
}
