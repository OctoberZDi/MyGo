package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/:name/:id", func(context *gin.Context) {
		var uri URI
		if err := context.ShouldBindUri(&uri); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, gin.H{"name": uri.Name, "id": uri.ID})
	})
	err := router.Run(":8008")
	if err == nil {
		return
	}
}
