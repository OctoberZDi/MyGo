package main

import (
	"Example/entity"
	hw "Example/hello-world"
)

func main() {
	hw.SayHello("tom")
	hw.SayBye()

	user := entity.User{Name: "zhangdi", Age: 20, Score: 20}
	// 首字母大写的才能调用
	entity.GetUserInfo(user)
}
