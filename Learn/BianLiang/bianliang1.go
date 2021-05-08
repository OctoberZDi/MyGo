package main

import "fmt"

// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
// var 语句可以出现在包或函数级别
func main() {
	var i int
	var c, python, java = true, false, "no!"
	fmt.Printf("i= %d c= %t python= %t java= %s ", i, c, python, java)

	// 短变量声明 简洁赋值语句 := 可在类型明确的地方代替 var 声明。
	name := "Shanghai"
	println("")
	print(name)
}
