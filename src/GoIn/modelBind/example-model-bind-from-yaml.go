package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// main 包中的不同的文件的代码不能相互调用，其他包可以。
// 选中引用的包，然后右键运行即可。
func main() {
	router := gin.Default()
	router.POST("loginYaml", func(context *gin.Context) {
		var loginJson Login

		// shouldBindYaml
		if err := context.ShouldBindYAML(&loginJson); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if loginJson.Password != "123456" || loginJson.User != "zhangdi" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in!"})
	})
	err := router.Run(":8003")
	if err == nil {
		return
	}
}
