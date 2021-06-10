package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(context *gin.Context) {
		firstName := context.DefaultQuery("firstName", "Zhangdi")
		lastName := context.Query("lastName")
		context.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
