package main

import (
	"Example/entity"
	hw "Example/hello-world"
)

func main() {
	hw.SayHello("tom")
	hw.SayBye()

	user := entity.User{Name: "zhangdi", Age: 20, Score: 20}
	entity.GetUserInfo(user)
}
