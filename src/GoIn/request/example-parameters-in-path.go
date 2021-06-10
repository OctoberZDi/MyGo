package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 路径匹配
func main() {
	student := Student{Name: "张迪", Score: 101}
	println(GetStudentInfo(student))

	router := gin.Default()

	// 路径中的参数
	// 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		message := name + " is " + action
		context.String(http.StatusOK, message)
		context.String(http.StatusOK, message+" "+strconv.FormatBool(context.FullPath() == "/user/:name/*action"))
	})

	router.POST("/user/groups", func(context *gin.Context) {
		context.String(http.StatusOK, "The available groups are [...]", "test...")
	})
	router.POST("/somePost", func(context *gin.Context) {

	})
	router.PUT("/somePut", func(context *gin.Context) {

	})
	router.DELETE("/someDelete", func(context *gin.Context) {

	})
	router.PATCH("/somePatch", func(context *gin.Context) {

	})
	router.HEAD("/someHead", func(context *gin.Context) {

	})
	router.OPTIONS("/someOptions", func(context *gin.Context) {

	})

	// listen and server on 0.0.0.0:8080(for windows "localhost:8080")
	err := router.Run(":3333")
	if err != nil {
		return
	}
}
