package main

import (
	"fmt"
	"os"
)

func main() {
	var numbers []int
	if numbers == nil {
		fmt.Println("======切片为空")
	} else {
		fmt.Println("======切片不为空")
	}

	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("length=%v \n", len(slice1))
	fmt.Printf("cap=%v \n", cap(slice1))
	for i, i2 := range slice1 {
		fmt.Printf("index=%d,value=%v ", i, i2)
	}
	fmt.Println("")

	// 获取前3个值
	slice11 := slice1[:3]
	fmt.Printf("%v \n", slice11)

	// 获取1~3的元素
	slice12 := slice11[1:3]
	fmt.Printf("%v \n", slice12)

	// 使用make初始化切片
	fmt.Println("使用make初始化切片")
	slice2 := make([]int, 10)
	if slice2 == nil {
		fmt.Println("======切片为空")
	} else {
		fmt.Println("======切片不为空")
	}

	fmt.Printf("length=%v \n", len(slice2))
	fmt.Printf("cap=%v \n", cap(slice2))

	fmt.Println("追加元素,追加到切片的末尾")
	slice2 = append(slice2, 100)
	slice2 = append(slice2, 200)
	slice2 = append(slice2, 300)
	fmt.Printf("length=%v \n", len(slice2))
	fmt.Printf("cap=%v \n", cap(slice2))
	if slice2 == nil {
		fmt.Println("======切片为空")
	} else {
		fmt.Println("======切片不为空")
	}

	// 在数组上使用range将传入index和值两个变量。
	// 上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for _, i2 := range slice2 {
		fmt.Printf("%v、", i2)
	}

	// 切片截取
	fmt.Println("切片截取")
	var slice22 = make([]int, 10)
	var slice23 = make([]int, 10)
	if slice2 != nil {
		// 从索引9（包含） 到索引12（不包含）
		slice22 = slice2[9:12]
		slice23 = slice2[9:13]
	}

	fmt.Println("slice2[9:12]:", slice22, "length:", len(slice22), "cap:", cap(slice22))
	fmt.Println("slice2[9:13]:", slice23, "length:", len(slice23), "cap:", cap(slice23))

	s := []struct{ x, y int }{{1, 11}, {2, 22}, {3, 33}}
	if s != nil {
		for index, ele := range s {
			fmt.Printf("index=%d,ele=%v,ele.x=%v,ele.y=%v \n", index, ele, ele.x, ele.y)
		}
	}

	fmt.Println("测试cap不足的时候，容量变化")

	// nil 切片
	var s1 []int
	fmt.Println(s1, cap(s1), len(s1))
	if s1 == nil {
		fmt.Println("s is nil!")
	}

	// 简单循环range
	sum := 0
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		sum += num
	}
	fmt.Printf("sum=%v", sum)

	fmt.Printf("%v", os.Args)
}

func printSlice1(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
