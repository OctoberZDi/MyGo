package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	var router = gin.Default()

	// Set a lower memory limit for multipart forms (default is 32Mib)
	router.MaxMultipartMemory = 8 << 20
	router.POST("/multiUpload", func(context *gin.Context) {
		name := context.PostForm("name")
		email := context.PostForm("email")
		// Multipart form
		form, error1 := context.MultipartForm()
		if error1 != nil {
			log.Println("Error is occurs...", error1)
			return
		}
		if form == nil {
			log.Println("Files not exists...")
			return
		}
		// 获取文件
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			//file.Filename == "aa.txt"

			fNames := strings.Split(file.Filename, string('.'))
			err := context.SaveUploadedFile(file, "D:/logs/"+fNames[0]+"_back."+fNames[1])
			if err != nil {
				log.Println("Save error is occurs...", err)
			}
		}
		context.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("%d files uploaded!", len(files)), "name": name, "email": email})
	})
	err := router.Run(":8888")
	if err == nil {
		return
	}
}
