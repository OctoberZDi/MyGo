package main

import (
	"fmt"
	"time"
)

// https://www.yiibai.com/go/golang-timeouts.html#article-start
/*超时对于连接到外部资源或在不需要绑定执行时间的程序很重要。
在Go编程中由于使用了通道和选择(select),实现超时是容易和优雅的。
在这个示例中，假设正在执行一个外部调用，2秒后在通道c1上返回其结果。
这里是 select 实现超时。 res：= <-c1等待结果和<-Time。
等待在超时1秒后发送一个值。 由于选择继续准备好第一个接收，如果操作超过允许的1秒，则将按超时情况处理。
如果允许更长的超时，如：3s，那么从c2的接收将成功，这里将会打印结果。
运行此程序显示第一个操作超时和第二个操作超时。
使用此选择超时模式需要通过通道传达结果。这是一个好主意，
因为其他重要的Go功能是基于渠道和Select。现在看看下面的两个例子：计时器和ticker。
*/
func main() {

	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan1 <- "result 1"
	}()

	select {
	case res1 := <-chan1:
		fmt.Println(res1)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "result 2"
	}()

	select {
	case res := <-chan2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
