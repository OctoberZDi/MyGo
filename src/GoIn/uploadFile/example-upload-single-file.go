package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8MiB
	router.POST("/singleUpload", func(context *gin.Context) {
		// single file
		file, _ := context.FormFile("file")
		// log
		log.Println(file.Filename, file.Size)

		// multiple
		// multiForm, error1 := context.MultipartForm()
		// files -> form 的name
		// 		Content-Disposition: form-data; name="files"; filename="aa.txt"
		// files := multiForm.File["files"]

		err := context.SaveUploadedFile(file, "D:/logs/aa1.txt")
		if err != nil {
			return
		}

		context.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
	})
	err := router.Run(":7777")
	if err == nil {
		return
	}
}
