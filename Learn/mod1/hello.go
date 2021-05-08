package main

import (
	"fmt"
)

// 切片
func main() {

	const name1 string = "我爱工作"
	println(name1, len(name1))

	/* 声明实际变量 */
	var a1 int = 20

	/* 声明指针变量 */
	var ip *int
	fmt.Printf("ptr 的值为 : %x\n", ip)
	/* 指针变量的存储地址 */
	ip = &a1

	fmt.Printf("a1 变量的地址是: %x\n", &a1)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	return
}
func Hello() {
	fmt.Println("hello go！")
	bye()
}

func bye() {
	fmt.Println("bye bye go!")
}
