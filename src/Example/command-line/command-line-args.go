package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println("index=", i, "arg=", arg)
	}
	fmt.Println("hello world!")
}
