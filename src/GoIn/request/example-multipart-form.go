package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message, _ := context.GetPostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")

		context.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": message,
			"nick":    nick,
		})
	})

	err := router.Run(":5555")
	if err != nil {
		return
	}
}
