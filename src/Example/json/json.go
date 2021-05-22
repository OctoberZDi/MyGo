package main

import (
	"encoding/json"
	"fmt"
)

// https://www.yiibai.com/go/golang-json.html#article-start
// Go提供对JSON编码和解码的内置支持，包括内置和自定义数据类型。

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	boolB, _ := json.Marshal(true)
	fmt.Println("bool:", string(boolB))

	intB, _ := json.Marshal(1)
	fmt.Println("int:", string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println("float:", string(fltB))

	strB, _ := json.Marshal("hello")
	fmt.Println("string:", string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println("array:", string(slcB))

	mapB := map[string]string{"apple": "1", "peach": "2"}
	mapD, _ := json.Marshal(mapB)
	fmt.Println("map:", string(mapD))

	response1 := &Response1{1, []string{"apple", "banana"}}
	response1B, _ := json.Marshal(response1)
	fmt.Println("response1:", string(response1B))
	response2 := &Response2{1, []string{"apple2", "banana2"}}
	response2B, _ := json.Marshal(response2)
	fmt.Println("response2:", string(response2B))
}
