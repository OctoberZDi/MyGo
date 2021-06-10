package main

import "github.com/gin-gonic/gin"

//通常，JSON 用它们的 unicode 实体替换特殊的 HTML 字符，
//例如< 变成 \u003c。如果您想按字面意思对此类字符进行编码，
// 则可以改用 PureJSON。此功能在 Go 1.6 及更低版本中不可用。

func main() {
	router := gin.Default()
	// Serves unicode entities
	router.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	router.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	err := router.Run(":8012")
	if err != nil {
		return
	}
}
