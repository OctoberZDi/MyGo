package main

import (
	"fmt"
	"time"
)

// go并发，go语言支持并发，只需要通过关键字go来开启goroutime即可
// go 运行时管理的轻量级线程
func main() {
	// 执行后打印的hello world没有固定的先后顺序，因为他们是连个goroutime在执行
	go say("world")
	go say("！！！")
	say("hello")
}

func say(content string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf(content+"%d \n", i)
	}
}
