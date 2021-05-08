package main

import "fmt"

// 切片的复制和追加
func main() {
	slFrom := []int{1, 2, 3, 4, 5}
	slTo := make([]int, 10)
	println(cap(slFrom))
	println(len(slFrom))
	println(cap(slTo))
	println(len(slTo))
	//
	copy(slTo, slFrom)
	for i, i2 := range slTo {
		fmt.Printf("index = %d,val = %v \n", i, i2)
	}

	slTo = append(slTo, 4, 5, 6, 7, 8)
	println("我又加了5各元素")
	for i, i2 := range slTo {
		fmt.Printf("index = %d,val = %v \n", i, i2)
	}

	s1 := "hello world"
	for i, i2 := range s1 {
		fmt.Printf("index = %d,val = %c \n", i, i2)
	}

	sHello := s1[:5]
	println(sHello)

	// 将字符串 "hello" 转换为 "cello"
	c := []byte(sHello)
	c[0] = 'c'
	s2 := string(c)
	println(s2)
}
