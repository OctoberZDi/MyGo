package main

import "fmt"

func main() {

	var i = 10
	fmt.Printf("An integer:%d,it's location in memory: %p", i, &i)

	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
