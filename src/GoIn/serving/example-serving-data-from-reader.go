package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

//从读取器提供数据
func main() {
	router := gin.Default()

	router.GET("/someDataFromRender", func(context *gin.Context) {
		resp, err := http.Get("https://10.200.72.95/images/color.png")
		if err != nil {
			log.Println("图片读取错误", err.Error())
			return
		}

		reader := resp.Body
		// 延迟关闭
		defer func(reader io.ReadCloser) {
			err := reader.Close()
			if err != nil {
				log.Println("reader close error...", err.Error())
				return
			}
		}(reader)

		// 以下操作相当于下载服务端的一个图片，该文件位于nginx下
		contentLength := resp.ContentLength
		contentType := resp.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment;filename="gopher.png"`,
		}

		context.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	err := router.Run(":8013")
	if err == nil {
		return
	}
}
