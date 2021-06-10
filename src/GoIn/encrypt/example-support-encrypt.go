package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/encrypt", func(context *gin.Context) {
		context.String(http.StatusOK, "Log content is encrypted!")
		// 执行报错
		// listen tcp :443: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
		log.Fatal(autotls.Run(engine, "www.baidu.com", "www.baidu.com"))
	})
	err := engine.Run(":3456")
	if err != nil {
		return
	}
}
