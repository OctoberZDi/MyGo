package main

import "fmt"

// range关键字用于for循环数组array，切片slice,通道channel，和集合map元素
//在数组和切片中返回元素的索引和索引值，在集合中返回key-value对
func main() {

	// 切片range
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("index = %d,value = %d ", index, value)
	}
	fmt.Println()
	// 数组range
	nums2 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	for i, i2 := range nums2 {
		fmt.Printf("index = %d,value = %.2f ", i, i2)
	}
	fmt.Println()

	var myMap = make(map[string]string)
	myMap["name"] = "张迪"
	myMap["age"] = "21"
	myMap["address"] = "山东省临沂市"
	for key, value := range myMap {
		fmt.Printf("key = %s,value = %s;", key, value)
	}
	fmt.Println()
	name, ok := myMap["name"]
	fmt.Printf("name = %s,ok = %t", name, ok)
}
