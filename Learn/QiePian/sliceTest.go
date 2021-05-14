package main

import "fmt"

// Go 语言切片是对数组的抽象。

// Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，
// Go 中提供了一种灵活，功能 强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

// 切片
func main() {
	// 数组转切片
	var myArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice1 = myArray[0:5]
	for i, i2 := range myArray {
		println(i, i2)
	}
	
	println("===")
	for i := range myArray {
		println(i, myArray[i])
	}
	println("===")
	for i := range mySlice1 {
		println(i, mySlice1[i])
	}
	println("===")
	for i, i2 := range mySlice1 {
		println(i, i2)
	}
	fmt.Println(myArray)
	fmt.Println(mySlice1)

	// 切片操作
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)
	// 打印从索引1（包含） 到索引4（不包含）
	fmt.Println("numbers[1:4]", numbers[1:4])
	// 默认下限为0
	fmt.Println("numbers[:4]", numbers[:4])
	numbers2 := numbers[0:2]
	printSlice(numbers2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	//创建一个切片 make
	println("====分割线")
	// make (type,length,capacity 可省略)
	mySlice2 := make([]int, 0)
	fmt.Println(mySlice2, len(mySlice2))
	mySlice3 := make([]int, 0)
	fmt.Println(mySlice3, len(mySlice3))
	mySlice2 = append(mySlice2, 1, 2, 3)
	mySlice3 = append(mySlice3, 4, 5, 6)
	fmt.Println(mySlice2, len(mySlice2))
	fmt.Println(mySlice3, len(mySlice3))

	// 将一个数组切片追加到另一个的末尾
	mySlice3 = append(mySlice3, mySlice2...)
	fmt.Println(mySlice3, len(mySlice3))

	// 合并多个slice到mySlice4
	mySlice4 := make([]int, 0)
	mySlice4 = append(mySlice4, 8, 8, 8)
	fmt.Println(mySlice4)

	// 合并mySlice2,mySlice3到mySlice4
	mySlice4 = append(append(mySlice4, mySlice2...), mySlice3...)
	fmt.Println("合并后：", mySlice4)
	// slice遍历
	for index, element := range mySlice3 {
		fmt.Printf("index = %d,element = %d", index, element)
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
