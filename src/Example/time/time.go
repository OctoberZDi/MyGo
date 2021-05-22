package main

import (
	"fmt"
	"time"
)

// https://www.yiibai.com/go/golang-time.html#article-start
func main() {

	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.Local)
	p(then)

	p("Year=", now.Year())
	p("Month=", now.Month())
	p("Day=", now.Day())
	p("Hour=", now.Hour())
	p("Minute=", now.Minute())
	p("Second=", now.Second())
	p("Nanosecond=", now.Nanosecond())
	p("Weekday=", now.Weekday())

	after := now.Add(time.Hour * 24)
	p("after=", after)

	f := func(x, y int) int { return x + y }
	println("10+20=", f(10, 20))
}
