package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
//POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//
//names[first]=thinkerou&names[second]=tianou

func main() {
	router := gin.Default()
	router.POST("/form-post", func(context *gin.Context) {
		ids := context.QueryMap("ids")
		var idsStr, _ = json.Marshal(ids)
		names := context.PostFormMap("names")
		var namesStr, _ = json.Marshal(names)

		context.JSON(http.StatusOK, gin.H{
			"ids":   string(idsStr),
			"names": string(namesStr),
		})
	})

	err := router.Run(":7777")
	if err == nil {
		return
	}
}
