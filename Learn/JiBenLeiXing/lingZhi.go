package main

import (
	"fmt"
	"runtime"
)

/*
没有明确初始值的变量声明会被赋予 零值
数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）*/
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("Go runs on...")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s\n", os)
	}
}
