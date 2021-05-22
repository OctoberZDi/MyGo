package main

import "fmt"

func main() {

	// 定义一个切片
	slice := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Println("slice[1:4]", slice[1:4]) // ola
	fmt.Println("slice[:2]", slice[:2])   // go
	fmt.Println("slice[2:]", slice[2:])   // lang
	fmt.Println("slice[:]", slice[:])     // golang

	for i, s := range slice {
		fmt.Printf("Slice at %d is: %v\n", i, s)
	}

	// 定义一个数组
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range array1 {
		i *= 2
	}

	fmt.Println(len(array1[2:2]), len(array1[2:3]))

	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d, %v \n", len(slice1), slice1)
	}
}
