package main

import "fmt"

func main() {
	message := make(chan string, 3)
	go func(content string, message chan string) {
		message <- content
		fmt.Printf("say %v \n", content)
		//close(message)
	}("333", message)
	go sayHello("222", message)
	sayHello("111", message)

	// 输出通道信息
	println("输出通道长度和容量：", "len=", len(message), "cap=", cap(message))
	fmt.Println("1", <-message)
	fmt.Println("2", <-message)
	fmt.Println("3", <-message)
	fmt.Println("4", <-message)
}

func sayHello(content string, message chan string) {
	message <- content
	//close(message)
	fmt.Printf("say %v \n", content)
}
