package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义中间件

func main() {
	router := gin.Default()

	// 使用自定义中间件
	router.Use(Logger())

	router.GET("/test", func(context *gin.Context) {
		//MustGet 返回给定键的值（如果存在），否则它会发生恐慌。
		example := context.MustGet("example").(string)
		name := context.MustGet("name")
		log.Println("name:", name)
		log.Println("example:", example)
	})
	err := router.Run(":8010")
	if err != nil {
		return
	}
}

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		//set example variable
		context.Set("example", "12345")
		context.Set("name", "张迪")

		//before request
		context.Next()

		//after request
		latency := time.Since(now)
		log.Println("latency:", latency)

		// access the status we are sending
		status := context.Writer.Status()
		log.Println("location:", "********* I AM HERE **********")
		log.Println("status:", status)
	}
}
