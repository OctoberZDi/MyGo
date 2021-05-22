package main

import (
	"fmt"
	"sort"
)

// map
func main() {
	myMap := make(map[string]string)
	myMap["name"] = "张迪"
	myMap["age"] = "30"
	myMap["age"] = "31"
	myMap["address"] = "山东省济南市"
	for s, s2 := range myMap {
		fmt.Printf("key = %s,value = %s \n", s, s2)
	}

	// 方式1：初始化map
	var myMap2 map[string]string
	if myMap2 == nil {
		fmt.Println("myMap2 is nil!")
	}
	// 方式2：初始化map
	myMap1 := make(map[int]string)
	if myMap1 == nil {
		fmt.Println("myMap1 is nil!")
	}
	myMap1[1] = "张  迪"
	myMap1[2] = "朱  玉"
	myMap1[3] = "张筱满"
	fmt.Println("我新创建了一个map")
	for key, value := range myMap1 {
		fmt.Printf("key=%d,value=%v \n", key, value)
		if key == 3 {
			myMap1[3] = "aaaaa"
			fmt.Printf("修改后: \nkey=%d,value=%v \n", key, myMap1[key])
		}
	}
	// 判断key是否存在 双赋值检测
	_, isCommentExist := myMap["comment"]
	fmt.Println(isCommentExist)
	_, isAddressExist := myMap["address"]
	fmt.Println(isAddressExist)
	if _, ok2 := myMap["name"]; ok2 {
		fmt.Println(ok2)
	}
	if age, isAgeExist := myMap["age"]; isAgeExist {
		fmt.Printf("isAgeExist=%v,age=%v \n", isAgeExist, age)
	}

	_, isAge := myMap["age"]
	fmt.Println("是否存在age", isAge)
	delete(myMap, "age")
	_, isAge1 := myMap["age"]
	fmt.Println("是否存在age", isAge1)

	// 排序
	var (
		barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
			"delta": 87, "echo": 56, "foxtrot": 12,
			"golf": 34, "hotel": 16, "indio": 87,
			"juliet": 65, "kili": 43, "lima": 98}
	)
	fmt.Println("开始排序")
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v || ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v || ", k, barVal[k])
	}
}
