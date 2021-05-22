package main

import (
	"fmt"
	"os"
)

// defer用于确保稍后在程序执行中执行函数调用，常用于清理目的

// 创建一个文件，写入内容，在完成后关闭，可以使用defer处理

func main() {
	file := createFile("defer.txt")
	defer closeFile(file)
	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("creating file!")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing file！")
	// 写入内容1
	f.WriteString("Hello world,hello go,hello defer!!oh yes")
	// 写入内容2
	fmt.Fprintln(f, "33333333333333")
}

func closeFile(f *os.File) {
	fmt.Println("closing file！")
	f.Close()
}
