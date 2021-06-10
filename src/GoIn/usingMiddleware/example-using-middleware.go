package main

import (
	"github.com/gin-gonic/gin"
)

//使用中间件
func main() {
	// 默认情况下创建一个没有任何中间件的路由器
	r := gin.New()

	// 全局中间件
	// 即使您设置了 GIN_MODE=release，Logger 中间件也会将日志写入 gin.DefaultWriter。
	// 默认情况下 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//每个路由中间件，您可以添加任意数量的中间件。
	r.GET("benchmark", func(context *gin.Context) {
		// empty
	}, func(context *gin.Context) {
		// empty
	})

	// 授权组 AuthRequired() 路由中间件
	// authorized := r.Group("/", AuthRequired())
	// 完全一样：
	authorized := r.Group("/")
	// 每组中间件！在这种情况下，我们使用自定义创建
	// AuthRequired() 中间件就在“已授权”组中。
	authorized.Use(func(context *gin.Context) {
		// 多个路由封装进一个路由中,路由封装到一块使用中间件，但是在此后在使用r的时候都默认注册了中间件
	})
	{
		authorized.POST("/login", func(context *gin.Context) {

		})
		authorized.POST("/submit", func(context *gin.Context) {

		})
		authorized.POST("/read", func(context *gin.Context) {

		})

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", func(context *gin.Context) {

		})
	}
}
