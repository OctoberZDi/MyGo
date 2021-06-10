package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	router.Delims("{[{", "}]}")
	// 设置模板函数
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("E:/IDEA_Workspace/MyGo/src/GoIn/serving/templates/function/raw.tmpl")
	//router.LoadHTMLFiles("../templates/function/raw.tmpl")
	router.GET("customFunc", func(context *gin.Context) {
		context.HTML(http.StatusOK, "raw.tmpl", gin.H{"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC)})
	})
	err := router.Run(":8013")
	if err != nil {
		return
	}
}

// 自定义函数
func formatAsDate(time time.Time) string {
	year, month, day := time.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}
