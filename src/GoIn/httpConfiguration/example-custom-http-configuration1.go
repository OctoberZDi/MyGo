package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/customConfig2", func(context *gin.Context) {
		context.String(http.StatusOK, "Connected 2 succeed!")
	})
	err := http.ListenAndServe(":8011", engine)
	if err != nil {
		return
	}
}
