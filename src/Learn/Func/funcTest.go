package main

import "fmt"

// 引用传递
var a int = 100
var b int = 200

//函数
func main() {
	var c, python, java = true, true, false
	var i, j int = 1, 2
	fmt.Println(c, python, java)
	fmt.Println(i, j)
	// 函数调用
	maxValue := max(22, 33)
	fmt.Printf("10和33的最大值是： %d\n", maxValue)

	// 多返回值函数使用
	x, y := swap("world", "hello")
	fmt.Printf("输入两个参数，交换后的内容：%s %s", x, y)

	fmt.Println()

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	swapEx(&a, &b)
	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

// 定义一个函数：返回两个值参数的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 定义一个函数，交换两个值的位置
func swap(x, y string) (string, string) {
	return y, x
}

func swapEx(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
