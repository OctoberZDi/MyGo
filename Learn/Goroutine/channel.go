package main

import "fmt"

// https://www.runoob.com/go/go-concurrent.html
// 通道channel用来传递数据的一个数据结构
// 通道可用于两个goroutine之间通过传递一个指定类型的值来同步运行和通讯。
// 操作符<-用于指定通道方向，发送或接收，如果未指定方向，则为双向通道。
// 例如
// ch <- v 把v发送到通道ch
// v:=<-ch // 从ch接收数据，并把值赋给v

func main() {

	// 声明一个通道 使用chan关键字，通道在使用前必须先创建 make创建的不会为nil
	ch := make(chan int)
	// 以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
	s := []int{1, 2, 3, 4, 5, 6}
	go sum(s[:len(s)/2], ch, "x")
	go sum(s, ch, "y")
	go sum(s[len(s)/2:], ch, "z")

	// 由于使用go程调用的sum方法，所以导致执行顺序不一致，xyz的值也不一定，按go程执行的先后顺序来决定其值
	x := <-ch
	y := <-ch
	z := <-ch
	fmt.Printf("x=%v \n", x)
	fmt.Printf("y=%v \n", y)
	fmt.Printf("z=%v \n", z)
	fmt.Printf("x+y+z=%v \n", x+y+z)
}

func sum(s []int, c chan int, t string) {
	sum := 0
	for _, val := range s {
		sum += val
	}

	// 把sum发送到通道c
	c <- sum
	fmt.Printf("t=%v,s=%v,sum=%v \n", t, s, sum)
}
