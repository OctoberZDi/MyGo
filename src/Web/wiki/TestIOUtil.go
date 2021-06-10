package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {

	//测试code
	r := gin.Default()
	println(r)
	err := ioutil.WriteFile("test.txt", []byte("你好，世界！"), 0666)
	if err != nil {
		return
	}

	fileContent, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Printf("读取文件内容失败！")
	}
	fmt.Printf("文件内容： %s", fileContent)

	fs, err := ioutil.ReadDir("D:/cg")
	for index, info := range fs {
		fmt.Printf("index= %d,fileName = %s,fileMode = %d, modeTime = %s,isDir = %t,size = %d, sys = %s;", index, info.Name(), info.Mode(), info.ModTime(), info.IsDir(), info.Size(), info.Sys())
		fmt.Println()
	}
}
