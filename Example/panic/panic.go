package main

import "os"

func main() {

	v, err := os.Create("./1111")
	println(v)
	if err != nil {
		panic(err)
	}
	panic("a problem")
}
