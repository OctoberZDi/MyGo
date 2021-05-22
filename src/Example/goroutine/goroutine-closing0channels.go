package main

import "fmt"

// https://www.yiibai.com/go/golang-closing-channels.html#article-start
// 关闭通道表示不会再发送更多值。
/*在这个例子中，我们将使用一个作业通道来完成从main()goroutine到worker goroutine的工作。
当没有更多的工作时，则将关闭工作通道。
这里是工作程序goroutine。 它反复从j的工作接收more := <-jobs。
在这种特殊的2值形式的接收中，如果作业已关闭并且已经接收到通道中的所有值，
则 more 的值将为 false。当已经完成了所有的工作时，使用这个通知。
这会通过作业通道向工作线程发送3个作业，然后关闭它。*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			fmt.Println(j, more)
			if more {
				fmt.Println("received job ", j)
			} else {
				fmt.Println("received all jobs ")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("sent job ", i)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	fmt.Println(<-done)
}
