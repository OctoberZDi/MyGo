package main

import (
	"fmt"
	"src/Model/entitys"
)

func main() {

	animal := entitys.Animal1{
		Name:     "Cat",
		Interest: "Sleeping",
	}

	fmt.Println(animal.Name, "likes", animal.Interest)
}
