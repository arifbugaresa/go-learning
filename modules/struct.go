package modules

import "fmt"

// Struct declaration
type student struct {
	name  string
	grade int
}

func LearningStruct() {
	StructLiteral()
	EmbeddedStruct()
	AnonymousStruct()
}

func StructLiteral() {

	// Implementation of struct
	var student1 student
	student1.name = "john doe"
	student1.grade = 2

	fmt.Println("name  :", student1.name)
	fmt.Println("grade :", student1.grade)

	// Implementation of struct
	var student2 = student{
		name:  "Jhon",
		grade: 3,
	}

	fmt.Println("name  :", student2.name)
	fmt.Println("grade :", student2.grade)

	// Implementation of struct
	var student3 = student{}
	student3.name = "wick"
	student3.grade = 2

	fmt.Println("name  :", student3.name)

}

type studentPerson struct {
	grade int
	person
}

type person struct {
	name string
	age  int
}

func EmbeddedStruct() {
	var doeData = person{
		name: "doe",
		age:  21,
	}

	var doe = studentPerson{
		person: doeData,
		grade:  2,
	}

	fmt.Println("name  :", doe.name)
	fmt.Println("age   :", doe.age)
	fmt.Println("grade :", doe.grade)
}

func AnonymousStruct() {
	var john = struct {
		person
		grade int
	}{}

	john.person = person{"wick", 21}
	john.grade = 2

	fmt.Println("name  :", john.person.name)
	fmt.Println("age   :", john.person.age)
	fmt.Println("grade :", john.grade)
}
