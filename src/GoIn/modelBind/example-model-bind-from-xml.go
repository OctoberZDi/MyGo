package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Example for binding XML (
//	<?xml version="1.0" encoding="UTF-8"?>
//	<root>
//		<user>manu</user>
//		<password>123</password>
//	</root>)
func main() {
	router := gin.Default()
	router.POST("loginXml", func(context *gin.Context) {
		var loginJson Login

		if err := context.ShouldBindXML(&loginJson); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if loginJson.Password != "123456" || loginJson.User != "zhangdi" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": "you are logged in!"})
	})
	err := router.Run(":8002")
	if err == nil {
		return
	}
}
