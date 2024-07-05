package main

import "fmt"

// LearningVariable
// Various ways to create variables in Golang
func LearningVariable() {
	var a string = "hello"

	var b, c int = 1, 2

	var (
		d = true
		e int
	)

	f := "e"

	fmt.Println(a, b, c, d, e, f)
}
