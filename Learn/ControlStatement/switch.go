package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	dir, _ := os.UserHomeDir()
	dir1, _ := os.UserCacheDir()
	dir2, _ := os.UserConfigDir()
	fmt.Printf("%v\n%v\n%v\n", dir, dir1, dir2)
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	test1()
}

func test1() {
	now := time.Now()
	fmt.Println(now)
	fmt.Printf("%4d年%02d月%02d日", now.Year(), now.Month(), now.Day())
}
