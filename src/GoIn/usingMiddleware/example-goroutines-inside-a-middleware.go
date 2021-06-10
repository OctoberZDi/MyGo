package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// *********** 在中间件或处理程序中启动新的 Goroutines 时，不
// 应使用其中的原始上下文，而必须使用只读副本*****************
func main() {

	router := gin.Default()
	// 异步
	router.GET("longAsync", func(context *gin.Context) {
		// 使用副本
		ctxCopy := context.Copy()
		go func() {
			// 用time.Sleep()模拟一个长任务
			time.Sleep(5 * time.Second)
			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path:" + ctxCopy.Request.URL.Path)
		}()
	})
	// 同步
	router.GET("longSync", func(context *gin.Context) {
		// 用time.Sleep()模拟一个长任务
		time.Sleep(5 * time.Second)
		// 因为没有使用goRoutine，所以不用使用context副本
		log.Println("Done! in path:" + context.Request.URL.Path)
	})
	err := router.Run(":8010")

	if err != nil {
		return
	}
}
