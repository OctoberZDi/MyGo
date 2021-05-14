package main

import (
	"fmt"
	"time"
)

// https://www.runoob.com/go/go-concurrent.html
// 通道缓冲区
// 带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态
// 就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。
// 不过，由于缓冲区大小是有限的，所以还是必须由接收端来接收数据的，否则缓冲区一满，数据发送端就无法发送数据了
// make 第二个参数指定缓冲区大小
func main() {
	// 1、定义一个可以存储整数类型的带缓冲的通道，缓冲区大小为3
	ch := make(chan int, 3)
	// 因为ch是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)
	// 向一个已经关闭的信道发送数据会引发程序恐慌panic
	// ch <- 4 // panic: send on closed channel
	x, ok := <-ch
	fmt.Println("======x", x, ok)
	// 获取这两个数据
	fmt.Println(<-ch, "len=", len(ch), "cap=", cap(ch))
	fmt.Println(<-ch, "len=", len(ch), "cap=", cap(ch))
	fmt.Println(<-ch, "len=", len(ch), "cap=", cap(ch))
	y, ok1 := <-ch
	fmt.Println("======y", y, ok1)
	fmt.Println(<-ch, "len=", len(ch), "cap=", cap(ch))
	fmt.Println("演示channel的遍历")

	// 2、通道遍历与通道关闭
	ch1 := make(chan int, 10)
	go fibonacci(cap(ch1), ch1)
	// range函数遍历每个通道中接收到的数据，因为c在发送玩10个数据之后就关闭了通道
	// 所以我们这里range函数在接收到10个数据之后就结束了
	// 如果上面的ch1通道不关闭，那么range函数就不会结束，从而在接收第11 个数据的时候就阻塞了
	for i := range ch1 {
		fmt.Print(i, " ")
	}

	// 3、channel是可以控制读写权限的
	ch2 := make(chan int, 2) // 可读可写
	ch2 <- 21
	ch2 <- 22
	fmt.Println("len=", len(ch2), <-ch2, <-ch2)
	ch3 := make(chan<- int, 2) // 只写
	ch3 <- 31                  // ok
	ch3 <- 32                  // ok

	// x := <-ch3 // Invalid operation: <-ch3 (receive from the send-only type chan<- int)
	fmt.Println("len=", len(ch3), cap(ch3))
	fmt.Println("==========")
	ch4 := make(<-chan int, 2) // 只读
	// x := <-ch4                 // ok,但是因为没有写入，直接读取会报错 fatal error: all goroutines are asleep - deadlock!
	// ch4 <- 41 // Invalid operation: ch4 <- 41 (send to the receive-only type <-chan int)
	fmt.Println("len=", len(ch4), cap(ch4))
}

// 遍历
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i <= n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(1000 * time.Millisecond)
	}

	// 关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
	close(c)
	fmt.Println("")
}
