package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	age   int
	class int
	grade float32
}

func main() {

	var student1 student
	var student2 student

	student1.name = "sandeep"
	student1.id = 1
	student1.age = 20
	student1.class = 10
	student1.grade = 85.5
	student2.name = "sandeep"
	student2.id = 2
	student2.age = 20
	student2.class = 10
	student2.grade = 85.5
	fmt.Println(student1)
	fmt.Println(student2)

	fmt.Println("name", student1.name)
	fmt.Println("id", student1.id)
	fmt.Println("age", student1.age)
	fmt.Println("class", student1.class)
	fmt.Println("grade", student1.grade)
	fmt.Println("name", student2.name)
	fmt.Println("id", student2.id)
	fmt.Println("age", student2.age)
	fmt.Println("class", student2.class)
	fmt.Println("grade", student2.grade)

}
