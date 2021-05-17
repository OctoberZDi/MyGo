package main

import "fmt"

func main() {
	i := 1
	for i < 3 {
		fmt.Printf("i = %d \n", i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d \n", i)
	}
	for {
		fmt.Println("loop")
		break
	}
	for {
		fmt.Println("loop")
		break
	}

}
