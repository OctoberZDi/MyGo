package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := setRouter()
	err := router.Run(":8004")
	if err != nil {
		return
	}
}

func setRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})
	return engine
}
