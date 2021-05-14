package main

import (
	"fmt"
	"time"
)

func fibonacci1(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c, " ")
			time.Sleep(1000 * time.Millisecond)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)
}
