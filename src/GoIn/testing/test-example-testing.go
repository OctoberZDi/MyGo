package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := setRouter1()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "Hello world!", recorder.Body.String())
}

func setRouter1() *gin.Engine {
	engine := gin.Default()
	engine.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello world!")
	})
	err := engine.Run(":8004")
	if err != nil {
		return nil
	}
	return engine
}
