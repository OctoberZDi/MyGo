package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	server := http.Server{
		Addr:           ":8011",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	engine.GET("/customConfig", func(context *gin.Context) {
		context.String(http.StatusOK, "Connected succeed!")
	})
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
