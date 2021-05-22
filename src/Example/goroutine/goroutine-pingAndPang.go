package main

import "fmt"

// https://www.yiibai.com/go/golang-channel-directions.html
/*当使用通道作为函数参数时，可以指定通道是否仅用于发送或接收值。
这种特殊性增加了程序的类型安全性。
此ping功能只接受用于发送值的通道。尝试在此频道上接收将是一个编译时错误。
ping函数接受一个通道接收(ping)，一个接收发送(ping)。*/
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) {
	// pings只可接收
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// 只可发送
	msg := <-pings
	pongs <- msg
}
