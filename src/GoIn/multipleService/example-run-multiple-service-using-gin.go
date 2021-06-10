package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"time"
)

var group errgroup.Group

func router01() http.Handler {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.GET("/server01", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Welcome server01!"})
	})
	return engine
}

func router02() http.Handler {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.GET("/server02", func(context *gin.Context) {
		context.JSONP(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Welcome server02!"})
	})
	return engine
}

// Run multiple service using gin
// 使用gin运行多服务
func main() {

	server01 := http.Server{Addr: ":8001", Handler: router01(), ReadTimeout: 5 * time.Second, WriteTimeout: 5 * time.Second}
	server02 := http.Server{Addr: ":8002", Handler: router02(), ReadTimeout: 5 * time.Second, WriteTimeout: 5 * time.Second}

	// 异步启动server01
	group.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	// 异步启动server02
	group.Go(func() error {
		err := server02.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	// Wait 阻塞，直到所有来自 Go 方法的函数调用都返回，然后
	// 从它们返回第一个非零错误（如果有）。
	err := group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
