package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie("gin-cookie")
		if err != nil {
			// 没有取到cookie值
			cookie = "NotSet"
			context.SetCookie("gin-cookie", "testValue", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value:%s\n\n", cookie)
	})

	engine.GET("/clearCookie", func(context *gin.Context) {
		cookie, err := context.Cookie("gin-cookie")
		if err == nil {
			// 代表可以取到cookie值
			context.SetCookie("gin-cookie", "clearedValue", 3600, "/", "localhost", false, true)
		}

		fmt.Printf("Cookie value:%s\n\n", cookie)
	})
	err := engine.Run(":8003")
	if err != nil {
		return
	}
}
