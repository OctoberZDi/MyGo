package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟从form表单中绑定属性

func main() {
	router := gin.Default()
	router.POST("/loginForm", func(context *gin.Context) {
		var form Login
		if err := context.ShouldBind(&form); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "zhangdi" || form.Password != "123456" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in!"})
	})

	err := router.Run(":8004")
	if err == nil {
		return
	}
}
