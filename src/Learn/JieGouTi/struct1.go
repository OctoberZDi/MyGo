package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 创建一个结构体 key=>value格式
	stu1 := Student{Name: "ZhangDi", Age: 30, Score: 100}
	fmt.Printf("Name:%s,Age:%d,Score:%d", stu1.Name, stu1.Age, stu1.Score)

	fmt.Println("")
	// Age,Score被隐式赋值
	stu2 := Student{Name: "XiaoMan"}
	fmt.Printf("Name:%s,Age:%d,Score:%d", stu2.Name, stu2.Age, stu2.Score)

	// 创建一个结构体stu3 不使用变量名
	stu3 := Student{"ZhuYu", 28, 122}
	fmt.Println("")
	fmt.Printf("Name:%s,Age:%d,Score:%d", stu3.Name, stu3.Age, stu3.Score)

	var stu4 Student
	stu4.Name = "Tom"
	stu4.Age = 10
	stu4.Score = 100
	fmt.Println("")
	fmt.Printf("Name:%s,Age:%d,Score:%d", stu4.Name, stu4.Age, stu4.Score)

	fmt.Println("")
	if result, err := json.Marshal(stu1); err == nil {
		fmt.Println(string(result))
	}
}

//结构体中属性的首字母大小写问题

// Student 首字母大写相当于 public。
// 首字母小写相当于 private。
// 当要将结构体对象转化为json的时候，对象中的属性首字母必须是大写，才能正常转化为JSON
type Student struct {
	// `json:"xxx"` 标识使用tag标记要返回的字段名
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}
