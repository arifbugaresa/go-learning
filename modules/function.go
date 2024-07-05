package modules

import "fmt"

// LearningFunction
// A function is a block of code created within a program to be used repeatedly.
func LearningFunction() {
	FunctionParameter()
	FunctionReturnValue()
	FunctionPredefinedReturnValue()
	FunctionReturnMultipleValue()
	FunctionVariadic()
	FunctionClosure()
}

func FunctionParameter() {
	printAngka(1, 2)
}

func printAngka(angka1 int, angka2 int) {
	fmt.Println("angka pertama", angka1)
	fmt.Println("angka kedua", angka2)
}

func FunctionReturnValue() {
	result := introduction("Doe")
	fmt.Println(result)
}

func introduction(name string) string {
	return "Hello My Name Is " + name
}

func FunctionPredefinedReturnValue() {
	result := tambahAngka(4, 5)
	fmt.Println(result)
}

func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

func FunctionReturnMultipleValue() {
	firstName, lastName := introduction2("John", "Doe")
	fmt.Println(firstName, lastName)
}

func introduction2(firstName string, lastName string) (string, string) {
	introFirstName := "Hello My First Name Is " + firstName
	introLastName := "Hello My Last Name Is " + lastName
	return introFirstName, introLastName
}

func FunctionVariadic() {
	var total = sum(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	fmt.Println(total)

	// function variadic parameter slice
	var numbers = []int{2, 6, 7, 8, 9, 10}
	var total2 = sum(numbers...)
	fmt.Println(total2)
}

func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func FunctionClosure() {
	var getMinMax = func(n []int) (int, int) {
		var minNum, maxNum int
		for i, e := range n {
			switch {
			case i == 0:
				maxNum, minNum = e, e
			case e > maxNum:
				maxNum = e
			case e < minNum:
				minNum = e
			}
		}
		return minNum, maxNum
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}

	var minNum, maxNum = getMinMax(numbers)
	fmt.Println("data :", numbers)
	fmt.Println("minNum :", minNum)
	fmt.Println("minNum :", maxNum)
}
