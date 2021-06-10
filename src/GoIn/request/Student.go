package main

import "strconv"

type Student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func GetStudentInfo(student Student) string {
	return student.Name + " " + strconv.Itoa(student.Score)
}
