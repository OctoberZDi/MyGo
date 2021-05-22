package main

import "fmt"

func main() {
	/*	a := 12
		doSomething1(&a)
		doSomething2(a)

		doSomething3(222)
		doSomething3(222, "hello")
		doSomething3(222, "hello", "world")*/
}
func doSomething1(a *int) {
	b := a
	println("do something ...", b)
}
func doSomething2(a int) {
	b := &a
	b2 := a
	println("do something ...", b)
	println("do something ...", b2)
}

func doSomething3(a int, who ...string) {
	fmt.Printf("a =  %d", a)
	fmt.Println()
	fmt.Printf("who = %v\n,length = %d", who, len(who))
}
