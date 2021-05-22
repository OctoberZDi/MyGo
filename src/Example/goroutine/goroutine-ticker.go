package main

import (
	"fmt"
	"time"
)
// https://www.yiibai.com/go/golang-tickers.html#article-start
// ticker用于定期做一些事情，周期性的执行指导停止
// NewTicker返回一个新的Ticker，其中包含将发送的频道
// 每次滴答之后通道上的时间。滴答周期为
// 由duration参数指定。实时报价器将调整时间
// 间隔或滴答滴答，以弥补接收速度慢的问题。
// 持续时间d必须大于零；如果没有，NewTicker将
// 恐慌。停止行情自动收录器以释放关联的资源。
func main() {
	//报价器使用与计时器类似的机制：
	//发送值的通道。在这里，我们将使用
	//通道内建的`range`进行迭代
	//这些值每500毫秒到达一次。
	ticker := time.NewTicker(time.Millisecond * 5000)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 20000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
