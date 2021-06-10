package main

import (
	"github.com/gin-gonic/gin"
)

// 绑定form的CheckBox
func main() {
	router := gin.Default()
	router.Any("/bindCheckbox", formHandler)

	err := router.Run(":8010")
	if err != nil {
		return
	}
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	err := c.ShouldBind(&fakeForm)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}
