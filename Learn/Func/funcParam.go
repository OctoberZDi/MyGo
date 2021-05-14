package main

import (
	"fmt"
)

type options struct {
	name    string
	age     int
	comment string
}

func main() {
	// 按值传递 call by value
	hello := "hello world,hello function!"
	sayHello1(hello)
	// 按值传递可以对副本的值进行更改，但不会影响到原来的值
	fmt.Println(hello)

	// 按引用传递 call by reference
	// 如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，
	// 比如 &variable）传递给函数，这就是按引用传递，
	// 几乎在任何情况下，传递指针（一个 32 位或者 64 位的值）的消耗都比传递副本来得少。
	//
	// 在函数调用时，像切片（slice）、字典（map）、接口（interface）、
	// 通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。
	sayHello2(&hello)

	fmt.Println(hello)

	// 多个参数返回
	fmt.Println(swap22(100, 200))

	// 可变长参数
	greeting("hello", "Tom", "Jimmy")
	names := []string{"Kitty", "Jerry", "Jack"}
	greeting("goodbye", names...)

	// 使用结构
	dealStruct("hello world", options{name: "zhangdi", age: 29,comment: "bye bye!!!"})
}

func sayHello1(content string) {
	b := content + "&&"
	fmt.Println(b)
}

func sayHello2(content *string) {
	//
	*content = "hello world,hello function! ！！"
	fmt.Println(*content)
}

func swap22(x, y int) (int, int, int) {
	return x, y, 300
}

func greeting(prefix string, who ...string) {
	for _, s := range who {
		fmt.Println(s)
	}
	fmt.Println(prefix, who)
}

// 使用
func dealStruct(prefix string, options2 options) {
	fmt.Println("处理结构体参数类型")
	fmt.Printf("%v,this is =%v,%d years old,comment is %v!", prefix, options2.name, options2.age, options2.comment)
}
