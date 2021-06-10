package main

import "fmt"

//定义一个接口 Simpler，它有一个 Get() 方法和一个 Set()，Get() 返回一个整型值，Set() 有一个整型参数。创建一个结构体类型 Simple 实现这个接口。
//接着定一个函数，它有一个 Simpler 类型的参数，调用参数的 Get() 和 Set() 方法。在 main 函数里调用这个函数，看看它是否可以正确运行。

type Simpler interface {
	Get() int
	Set(a int)
}

type Simple struct {
	value int
}

func (s Simple) Get() int {
	return s.value
}

func (s *Simple) Set(a int) {
	s.value = a
	//s.value = a
}

func displayValue(simpler Simpler) {
	fmt.Printf("get->%d\n", simpler.Get())
	simpler.Set(100)
	fmt.Println("set.......")
	fmt.Printf("get->%d\n", simpler.Get())
}
func main() {
	simple := &Simple{value: 88}
	displayValue(simple)
}
