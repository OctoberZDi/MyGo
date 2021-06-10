package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//分组路由

func login(context *gin.Context) http.HandlerFunc {
	fmt.Println("login..")
	return nil
}

func main() {
	router := gin.New()
	//router := gin.Default()

	// Simple group :v1
	v1 := router.Group("v1", func(context *gin.Context) {
		fmt.Println("v1中间件")
	})
	{
		v1.POST("login", func(context *gin.Context) {
			context.String(http.StatusOK, "login is ok!")
		})
		v1.POST("logout", func(context *gin.Context) {
			context.String(http.StatusOK, "logout is ok!")
		})
		v1.POST("submit", func(context *gin.Context) {
			context.String(http.StatusOK, "submit is ok!")
		})
	}

	// Simple group :v1
	v2 := router.Group("v2", func(context *gin.Context) {
		fmt.Println("v2中间件")
	})
	{
		v2.POST("read", func(context *gin.Context) {
			context.String(http.StatusOK, "read is ok!")
		})
		v2.POST("write", func(context *gin.Context) {
			context.String(http.StatusOK, "write is ok!")
		})
	}
	err := router.Run(":8877")
	if err == nil {
		return
	}
}
