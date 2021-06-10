package main

import "fmt"

// IAnimal 定义一个接口
type IAnimal interface {
	Eat()
	Run()
}

type Cat struct {
	name string
	sex  string
}

// Eat 实现Eat方法
func (cat Cat) Eat() {
	fmt.Printf("%s : %s -> is eatting! \n", cat.name, cat.sex)
}

// Run 实现 Run方法
func (cat Cat) Run() {
	fmt.Printf("%s : %s -> is running! \n", cat.name, cat.sex)
}

func main() {
	cat := Cat{
		name: "Tom",
		sex:  "boy",
	}

	cat.Eat()
	cat.Run()
}
