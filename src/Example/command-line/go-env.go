package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("myName", "张迪")
	os.Setenv("FOO", "1")
	fmt.Println("myName:", os.Getenv("myName"))
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	for index, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(index,  pair)
	}
}
