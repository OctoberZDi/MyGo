package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 参数查询
func main() {

	stu := Student{Name: "张迪", Score: 101}
	println(GetStudentInfo(stu))

	router := gin.Default()
	// http://localhost:8080/ping?name=zhangdi&password=123456
	//Querystring parameters
	router.GET("/ping", func(context *gin.Context) {
		name, ok := context.GetQuery("name")
		if !ok {
			name = "The name is unset!"
		}

		// shortcut for context.Request.URL.Query().Get("lastname")
		lastName := context.Request.URL.Query().Get("lastname")
		fmt.Println(lastName)
		password, ok := context.GetQuery("password")
		if !ok {
			password = "The password is unset!"
		}

		other, ok := context.GetQuery("other")
		if !ok {
			other = "Ex!"
		}
		context.JSON(200, gin.H{
			"message":  "pong",
			"name":     name,
			"password": password,
			"other":    other,
		})
	})
	// listen and server on 0.0.0.0:8080(for windows "localhost:8080")
	err := router.Run(":4444")
	if err != nil {
		return
	}
}
