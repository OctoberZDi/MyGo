package main

import (
	"fmt"
	"time"
)

//计算函数执行时间
func main() {
	testTime()
}

func testTime() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		println(i)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("循环执行了： %v", delta)
}
