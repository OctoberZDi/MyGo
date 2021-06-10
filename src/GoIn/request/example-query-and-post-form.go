package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Another example: query + post form
func main() {
	router := gin.Default()

	// query + post form
	router.POST("form-post", func(context *gin.Context) {
		id := context.Query("id")
		page := context.DefaultQuery("page", "0")
		name := context.PostForm("name")
		message := context.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		context.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
	err := router.Run(":6666")
	if err != nil {
		return
	}
}
