package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		var testHeader TestHeader
		err := context.ShouldBindHeader(&testHeader)
		// 上面的简写
		//err1 := context.ShouldBindWith(&testHeader, binding.Header)
		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			return
		}

		fmt.Printf("%#v\n", testHeader)
		context.JSON(http.StatusOK, gin.H{"Rate": testHeader.Rate, "Domain": testHeader.Domain})
	})
	err := router.Run(":8009")
	if err != nil {
		return
	}
}
