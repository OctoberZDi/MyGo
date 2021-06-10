package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//使用 AsciiJSON 生成带有转义非 ASCII 字符的纯 ASCII JSON。
func main() {
	router := gin.Default()

	router.GET("/asciiJson", func(context *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output:{"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		context.AsciiJSON(http.StatusOK, data)
	})
	err := router.Run(":8012")
	if err != nil {
		return
	}
}
