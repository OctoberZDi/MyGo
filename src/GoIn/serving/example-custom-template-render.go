package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//您也可以使用自己的 html 模板渲染

func main() {
	router := gin.Default()
	htmlTemp := template.Must(template.ParseFiles("E:/IDEA_Workspace/MyGo/src/GoIn/serving/templates/custom/index1.tmpl", "E:/IDEA_Workspace/MyGo/src/GoIn/serving/templates/custom/index2.tmpl"))
	router.SetHTMLTemplate(htmlTemp)

	router.GET("/customIndex1", func(context *gin.Context) {
		context.HTML(http.StatusOK, "custom/index1.tmpl", gin.H{"title": "张迪测试title-1"})
	})
	router.GET("/customIndex2", func(context *gin.Context) {
		context.HTML(http.StatusOK, "custom/index2.tmpl", gin.H{"title": "张迪测试title-2"})
	})
	err := router.Run(":8013")
	if err != nil {
		return
	}
}
