package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用 JSONP 从不同域中的服务器请求数据。如果存在查询参数回调，则向响应正文添加回调。

func main() {
	router := gin.Default()
	router.GET("/JSONP", func(context *gin.Context) {
		data := gin.H{"foo": "bar"}

		//callback is x
		// Will output:
		// x({"foo": "bar"});
		context.JSONP(http.StatusOK, data)
	})

	// Listen and server on 0.0.0:8013
	err := router.Run(":8012")
	if err == nil {
		return
	}
	// client
	// curl http://127.0.0.1:8080/JSONP?callback=x
}
