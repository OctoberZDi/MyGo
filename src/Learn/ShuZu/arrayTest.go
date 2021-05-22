package main

import "fmt"

var types = [5]string{"A", "B", "C", "D"}

func main() {
	types[4] = "E"
	for i, s := range types {
		fmt.Printf("index= %d,value= %s \n", i, s)
	}

	for index := range types {
		fmt.Printf("index = %d \n", index)
	}
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item index=", i, "is", a[i])
	}
}
