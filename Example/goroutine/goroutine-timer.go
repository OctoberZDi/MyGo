package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer类型代表一个事件。
	// 当计时器到期时，当前时间将在C上发送，
	// 除非计时器是由AfterFunc创建的。
	// 必须使用NewTimer或AfterFunc创建一个计时器。
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 计时器
	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
