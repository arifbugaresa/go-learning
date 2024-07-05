package modules

import "fmt"

// LearningConstant
// Constants are data types whose values cannot be changed
func LearningConstant() {
	const a string = "hello"

	const b, c int = 1, 2

	const (
		d     = true
		e int = 2
	)

	fmt.Println(a, b, c, d, e)
}
