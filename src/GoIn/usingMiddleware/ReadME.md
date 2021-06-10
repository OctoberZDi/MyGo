golang的net/http设计的一大特点就是特别容易构建中间件。
gin也提供了类似的中间件。需要注意的是中间件只对注册过的路由函数起作用。
对于分组路由，嵌套使用中间件，可以限定中间件的作用范围。中间件分为全局中间件，单个路由中间件和群组中间件。


定义一个中间件
func MiddleWare() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("before middleware")
        c.Set("request", "clinet_request")
        c.Next()
        fmt.Println("before middleware")
    }
}
