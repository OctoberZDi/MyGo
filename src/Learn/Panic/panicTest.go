package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	regexpStr := MustCompile("hello")
	fmt.Println(regexpStr)

	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
}

func MustCompile(str string) *regexp.Regexp {
	regexp, err := regexp.Compile(str)
	if err != nil {
		panic("regexp: Compile(" + strconv.Quote(str) + "): " + err.Error())
	}
	return regexp
}
