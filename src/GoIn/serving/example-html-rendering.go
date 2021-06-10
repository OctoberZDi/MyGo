package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

//加载html文件
func main() {
	router := gin.Default()

	// will output:C:\Go;E:\IDEA_Workspace\MyGo
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println(filepath.Join(os.Getenv("TMPL_DIR"), "*"))
	// 路径是啥？
	//router.LoadHTMLGlob()
	// 绝对路径？？使用GoPath或自定义环境变量？TEMP_DIR但是我不知道怎么设置
	router.LoadHTMLFiles("E:/IDEA_Workspace/MyGo/src/GoIn/serving/templates/posts/index.tmpl", "E:/IDEA_Workspace/MyGo/src/GoIn/serving/templates/users/index.tmpl")
	router.GET("posts/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	router.GET("users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})
	err := router.Run(":8013")
	if err != nil {
		return
	}
}
