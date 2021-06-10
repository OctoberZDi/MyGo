package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// 返回渲染的结果是 xml,json,yaml,protobuf的example
func main() {
	router := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	router.GET("/someJson", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name":    "zhangdi",
			"message": "hey some json!",
		})
	})

	router.GET("/moreJson", func(context *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"name"`
			Message string `json:"message"`
		}

		msg.Name = "zhangdi2"
		msg.Message = "hey more json!"
		context.JSON(http.StatusOK, msg)
	})

	// xml render
	router.GET("/someXml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"message": "hey xml!", "status": http.StatusOK})
	})

	// yaml render
	router.GET("/someYaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"message": "hey yaml!", "status": http.StatusOK})
	})

	// protobuf 原始缓冲区
	router.GET("/someProtoBuf", func(context *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"

		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{Label: &label, Reps: reps}
		// 注意数据在响应中变成了二进制数据
		// 将输出 protoexample.Test protobuf 序列化数据
		context.ProtoBuf(http.StatusOK, data)
	})
	err := router.Run(":8011")
	if err != nil {
		return
	}
}
