package modules

import "fmt"

// LearningOperator
// operator is a special character that represents an action.
func LearningOperator() {
	ArithmeticOperators()
	UnaryOperator()
	ComparisonOperators()
	LogicalOperators()
}

// ArithmeticOperators
// used for operations that involve calculations.
func ArithmeticOperators() {
	// addition operator
	sum := 8 + 3
	fmt.Println(sum) // result is 13

	// subtraction operator
	subtract := 8 - 3
	fmt.Println(subtract) // result is 5

	// multiplication operator
	multiply := 8 * 3
	fmt.Println(multiply) // result is 24

	// division operator
	divide := 8 / 4
	fmt.Println(divide) // result is 2

	// modulus operator
	modulus := 8 % 3
	fmt.Println(modulus) // result is 2

	// augmented assigment
	var number = 8
	number += 10
	fmt.Println(number) // 18

	var number2 = 5
	number2 += 5
	fmt.Println(number2) // 10

}

func UnaryOperator() {
	number := 8
	number++
	fmt.Println(number) // 9

	number2 := 5
	number2--
	fmt.Println(number2) // 4
}

// ComparisonOperators
// Comparison operators are used to determine the truth of a condition.
func ComparisonOperators() {
	var number = 8

	fmt.Println(number > 5)  // true
	fmt.Println(number < 5)  // false
	fmt.Println(number >= 5) // true
	fmt.Println(number <= 5) // false
	fmt.Println(number == 5) // false
	fmt.Println(number != 5) // true
}

// LogicalOperators
// These operators are used to determine the truth of data combinations.
func LogicalOperators() {
	var a = true
	var b = false
	var c = true
	var d = false

	fmt.Println(a && c) // true

	fmt.Println(a && b) // false

	fmt.Println(a || b) // true

	fmt.Println(b || d) // false

	fmt.Println(!b && !d) // true

	fmt.Println(!a || b) // false
}
