package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "zhangdi,hello"
	// 判断是否以前缀开头
	hasPrefix := strings.HasPrefix(name, "z")

	println(hasPrefix)
	// 判断是否以后缀结尾
	hasSuffix := strings.HasSuffix(name, "di")
	println(hasSuffix)

	// 字符串包含
	isContains := strings.Contains(name, "zhangdi")
	println(isContains)

	// 判断索引位置，或子串
	index1 := strings.Index(name, "z")
	println(index1)
	// 最后的位置
	lastIndex1 := strings.LastIndex(name, "l")
	println(lastIndex1)

	// 字符串替换
	newName := strings.Replace(name, "hello", "how are you!", -1)
	println(newName)

	// 字符串中字符出现的次数
	count := strings.Count(name, "l")
	fmt.Printf("字母 l 出现了 %d 次\n", count)

	// 修改字符串大小写
	lowerName := strings.ToLower(name)
	fmt.Printf("lower: %s", lowerName)
	upperName := strings.ToUpper(lowerName)
	println()
	fmt.Printf("upper: %s", upperName)

	// fields 利用一个或多个空白符号作为分隔符，返回一个slice
	str := "The quick brown fox jumps over the lazy dog"
	strSlice := strings.Fields(str)
	for _, val := range strSlice {
		fmt.Println(val)
	}
	strSlice2 := strings.Split(str, "|")
	fmt.Printf("Splitted in slice: %v\\n", strSlice2)
}
