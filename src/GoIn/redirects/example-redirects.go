package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//发出 HTTP 重定向很容易。支持内部和外部位置。
	router.GET("/redirectBaidu", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	//从 POST 发出 HTTP 重定向。
	router.GET("/redirectNginx", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://10.200.72.95")
	})

	router.GET("/test", func(context *gin.Context) {
		// 相对于nginx部署的测试html
		context.Request.URL.Path = "/go/example-model-bind-html-checkbox.html"
		router.HandleContext(context)
	})
	router.GET("/test2", func(context *gin.Context) {
		// 相对于nginx部署的测试html
		context.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
