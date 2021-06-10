package entity

import "fmt"

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func GetUserInfo(user User) {
	fmt.Printf("Name=%v,Age=%d,Score=%d", user.Name, user.Age, user.Score)
}

func getName(user User) string {
	return user.Name
}
