package main

import "fmt"

// https://www.yiibai.com/go/golang-range-over-channels.html#article-start
// for和range语句如何为基本数据结构提供迭代。
// 还而已使用此语法对通通道接收的值进行迭代
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for s := range queue {
		fmt.Println(s)
	}
}
