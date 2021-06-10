package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// StructA 使用自定义结构绑定表单数据请求
// form
type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
type StructD struct {
	NestedAnonyStruct struct {
		Fieldx string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(context *gin.Context) {
	var b StructB
	err := context.Bind(&b)
	if err != nil {
		log.Fatal(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}
func GetDataC(context *gin.Context) {
	var c StructC
	err := context.Bind(&c)
	if err != nil {
		log.Fatal(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"a": c.NestedStructPointer,
		"c": c.FieldC,
	})
}
func GetDataD(context *gin.Context) {
	var d StructD
	err := context.Bind(&d)
	if err != nil {
		log.Fatal(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonyStruct,
		"d": d.FieldD,
	})
}
func main() {
	engine := gin.Default()
	engine.GET("getB", GetDataB)
	engine.GET("getC", GetDataC)
	engine.GET("getD", GetDataD)
	err := engine.Run(":8003")

	if err != nil {
		return
	}
}
