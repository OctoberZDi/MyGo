package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Any("queryOrPost", func(context *gin.Context) {
		var person2 Person2
		// If `GET`, only `Form` binding engine (`query`) used.
		// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
		// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
		err := context.Bind(&person2)
		if err == nil {
			log.Println("====== Bind By Query String ======")
			log.Println(person2.Name)
			log.Println(person2.Address)
			log.Println(person2.Birthday)
			log.Println(person2.CreateTime)
			log.Println(person2.UnixTime)
		} else {
			log.Printf("***************err: %v\n", err.Error())
		}

		// 是不是json格式
		err1 := context.BindJSON(&person2)
		if err1 == nil {
			log.Println("====== Bind By Json ======")
			log.Println(person2.Name)
			log.Println(person2.Address)
			log.Println(person2.Birthday)
			log.Println(person2.CreateTime)
			log.Println(person2.UnixTime)
		} else {
			log.Printf("***************err1: %v\n", err1.Error())
		}
		context.JSON(http.StatusOK, gin.H{
			"Name":       person2.Name,
			"Address":    person2.Address,
			"Birthday":   person2.Birthday,
			"CreateTime": person2.CreateTime,
			"UnixTime":   person2.UnixTime,
		})
	})
	err := router.Run(":8007")
	if err != nil {
		return
	}
}
