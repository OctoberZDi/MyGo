package main

import "fmt"

func main() {
	fmt.Printf("The add function result is: %d ", add(100, 200))
	fmt.Println("")
	fmt.Printf("The add2 function result is: %d ", add(110, 220))
}

func add(x int, y int) int {
	return x + y
}

/**
当连续两个或多个函数的已命名形参类型相同时，
除最后一个类型外，其他都可以省略
*/
func add2(x, y int) int {
	return x + y
}
