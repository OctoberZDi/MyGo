package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 绑定查询字符串
//ShouldBindQuery 函数只绑定查询参数而不绑定发布数据。

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	err := route.Run(":8006")
	if err != nil {
		return
	}
}

func startPage(c *gin.Context) {
	var person Person
	// ShouldBindQuery 应该绑定form
	// BindQuery 必须绑定
	if c.BindQuery(&person) == nil {
		log.Println("====== Only Bind Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}
