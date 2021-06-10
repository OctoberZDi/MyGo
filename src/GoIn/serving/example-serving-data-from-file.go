package main

import (
	"github.com/gin-gonic/gin"
)

//从文件中提供数据
func main() {
	router := gin.Default()
	router.GET("/local/file", func(context *gin.Context) {
		// 返回文件内容
		context.File("D:/logs/aa.txt")
	})

	/*var fs http.FileSystem = "D:/logs/aaa.go"
	router.GET("/fs/file", func(context *gin.Context) {
		context.FileFromFS("D:/logs/file.go", fs)
	})*/
	err := router.Run(":8013")
	if err == nil {
		return
	}
}
