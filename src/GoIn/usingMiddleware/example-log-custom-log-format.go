package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	os "os"
	"time"
)

func main() {
	//禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	//gin.DisableConsoleColor()
	// 强制使用日志颜色
	gin.ForceConsoleColor()

	// Logging to a file.
	create, _ := os.Create("D:/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(create)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := gin.Default()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	err := router.Run(":8010")
	if err != nil {
		return
	}
}
