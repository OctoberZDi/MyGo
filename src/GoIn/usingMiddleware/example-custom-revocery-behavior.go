package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// 全局中间件
	// 即使您设置了 GIN_MODE=release，Logger 中间件也会将日志写入 gin.DefaultWriter。
	// 默认情况下 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// 恢复中间件从任何恐慌中恢复，如果有，则写入 500
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		// http.StatusInternalServerError 500
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	// Listen and serve on 0.0.0.0:8080
	err := r.Run(":8010")
	if err != nil {
		return
	}
}
