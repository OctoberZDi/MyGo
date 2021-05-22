package main

import "fmt"

func main() {
	// map
	myMap := make(map[string]string)
	myMap["name"] = "张迪"
	myMap["age"] = "28"
	myMap["score"] = "99"
	for key, value := range myMap {
		println(key, value)
	}

	name := "hello"
	for i := 0; i < len(name); i++ {
		//fmt.Println(i, string(name[i]))
		//fmt.Println(i, name[i])
	}

	for k, v := range name {
		fmt.Println(k, string(v))
	}
}
func Hello2() {
	fmt.Println("hello2 go！")
}

func bye2() {
	fmt.Println("bye bye go!")
}
