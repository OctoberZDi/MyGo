package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	where()
	f2()
	f()
}

func f2() {
	unnamed := func(name string) string { return "hello " + name + ",this is an anonymous function" }
	println(unnamed("world"))
}
func f() {
	for i := 0; i < 4; i++ {
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %T \n", g, g)
	}
}
