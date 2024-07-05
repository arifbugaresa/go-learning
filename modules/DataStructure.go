package modules

import "fmt"

// LearningDataStructure
// Data structures are collections of various variables grouped together to represent a specific entity or concept in programming.
func LearningDataStructure() {
	Array()
	Slice()
	Map()
}

func Array() {
	// One-dimensional Array
	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	names[2] = "Frank"
	names[3] = "Jack"

	// Accessing Elements
	fmt.Println(names[0], names[1], names[2], names[3])

	// Multidimensional Array
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// Accessing Elements
	fmt.Println("numbers1", numbers1[0][0])
	fmt.Println("numbers2", numbers2[0][1])
}

func Slice() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	// Accessing Elements
	fmt.Println(fruits[0]) // "apple"

	// Length of Slice
	fmt.Println(len(fruits)) // 4

	// Appending to a Slice:
	fruits = append(fruits, "orange")
	fmt.Println(fruits) // ["apple", "grape", "banana", "melon", "orange"]

	// Slicing a Slice:
	slice1 := fruits[1:3]
	fmt.Println(slice1) // ["grape", "banana"]

	// Copying a Slice:
	slice2 := make([]string, len(fruits))
	copy(slice2, fruits)

	fmt.Println(slice2) // ["apple", "grape", "banana", "melon", "orange"]

	// Iterating Over a Slice:

	for index, value := range fruits {
		fmt.Println(index, value)
	}
	// Output:
	// 0 apple
	// 1 grape
	// 2 banana
	// 3 melon
	// 4 orange
}

func Map() {
	var chicken map[string]int

	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("mei", chicken["mei"])         // mei 0

	// Deleting an Item from a Map
	delete(chicken, "januari")

	// Detecting the Existence of an Item with a Specific Key
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}
