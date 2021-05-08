package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		if sum > 4500 {
			fmt.Println("oho...")
		} else {
			fmt.Println("oh no...")
		}
		sum += i
	}
	fmt.Printf("The sum is: %v", sum)
}
