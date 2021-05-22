package main

import (
	"encoding/json"
	"fmt"
)

// Books 结构体
type Books struct {
	// tag 标记要返回的字段名。 json:xxx
	// 首字母大写相当于public
	Title string `json:"Title"`
	// 首字母小写相当于private
	Author  string `json:"author"`
	Subject string `json:"subject"`
	Id      int    `json:"id"`
}
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0
	v3 = Vertex{}      // X:0 Y:0被隐式地赋予
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	// 创建一个新的结构体 方式1
	book := Books{"Go语言", "张迪", "Go语言教程", 1002}
	// 创建一个新的结构体 方式2
	book2 := Books{Title: "Go语言", Author: "张迪", Subject: "Go语言教程2", Id: 1003}
	// 创建一个新的结构体 方式3 忽略的字段为0或空
	book3 := Books{Title: "Go语言", Author: "张迪", Subject: "Go语言教程23", Id: 1004}
	fmt.Println(book)
	fmt.Println(book2)
	fmt.Println(book3)

	println("=======")
	fmt.Println(&book.Title)
	fmt.Println(book.Title)
	// 结构体作为参数传递给函数
	fmt.Println("结构体作为参数传递给函数")
	println("原值")
	printBook(book)
	changeBook(book)
	println("第一次改变")
	printBook(book)
	// 结构体可以通过结构体指针来访问
	changeBook2(&book)
	println("第二次改变")
	printBook(book)

	// 结构体json操作 json.Marshal 将对象转换为json字符串
	if result, err := json.Marshal(&book); err == nil {
		fmt.Println(string(result))
	}
}

func changeBook(book Books) {
	book.Title = "Java语言"
}

func changeBook2(book *Books) {
	book.Title = "Python 语言"
}

func printBook(book Books) {
	fmt.Printf("Book Title : %s\n", book.Title)
	fmt.Printf("Book Author : %s\n", book.Author)
	fmt.Printf("Book Subject : %s\n", book.Subject)
	fmt.Printf("Book Id : %d\n", book.Id)
}
