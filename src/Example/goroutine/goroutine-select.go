package main

import (
	"fmt"
	"time"
)

// https://www.yiibai.com/go/golang-select.html#article-start
/* ******Go语言的选择(select)可等待多个通道操作。******

将goroutine和channel与select结合是Go语言的一个强大功能。
每个通道将在一段时间后开始接收值，以模拟阻塞在并发goroutines中执行的RPC操作。
我们将使用select同时等待这两个值，在每个值到达时打印它们。*/
func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		chan1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 10)
		chan2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("received", msg1)
		case msg2 := <-chan2:
			fmt.Println("received", msg2)
		}
	}
}
