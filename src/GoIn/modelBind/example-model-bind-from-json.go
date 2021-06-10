package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Example for binding JSON ({"user": "manu", "password": "123"})
func main() {
	router := gin.Default()
	router.POST("loginJson", func(context *gin.Context) {
		var loginJson Login

		if err := context.ShouldBindJSON(&loginJson); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if loginJson.Password != "123456" || loginJson.User != "zhangdi" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in!"})
	})
	err := router.Run(":8001")
	if err == nil {
		return
	}
}
