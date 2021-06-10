package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//提供静态文件
func main() {
	router := gin.Default()
	// Example:http://localhost:8013/assets/1.jpg
	router.Static("/assets", "D:/logs/images")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	err := router.Run(":8013")
	if err != nil {
		return
	}
}
