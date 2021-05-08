package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Printf("%4d年%02d月%02d日", now.Year(), now.Month(), now.Day())
	nowAdd := now.AddDate(1, 0, 0)
	fmt.Println(nowAdd)
}
