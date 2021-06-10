package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认在响应主体之前添加“while(1)”。
	// 如下请求结果： while(1);["lena", "austin", "foo"]
	router.GET("/secureJson", func(context *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		context.SecureJSON(http.StatusOK, names)
	})
	err := router.Run(":8012")
	if err != nil {
		return
	}
}
