package modules

import "fmt"

// car struct
type car struct {
	name string
}

// Desc method
func (c car) Desc() {
	fmt.Println(c.name)
}

func LearningMethod() {
	car1 := car{
		name: "BMW",
	}

	car1.Desc()
}
