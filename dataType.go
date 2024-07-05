package main

import "fmt"

// LearningDataType
// Various types of data types that exist in Golang
func LearningDataType() {
	// Types of numbers
	var (
		numeric1  uint8
		numeric2  uint16
		numeric3  uint32
		numeric4  uint64
		numeric5  int8
		numeric6  int16
		numeric7  int32
		numeric8  int64
		numeric9  int
		numeric10 rune
		numeric11 byte
	)

	fmt.Println(numeric1, numeric2, numeric3, numeric4, numeric5, numeric6, numeric7, numeric8, numeric9, numeric10, numeric11)

	// Types of decimals
	var (
		decimal1 float32
		decimal2 float64
	)

	fmt.Println(decimal1, decimal2)

	// Types of string
	var (
		string1 string
	)

	fmt.Println(string1)

	// Types of pointer
	var (
		pointer1 *string
		pointer2 *int
		pointer3 *float32
	)

	fmt.Println(pointer1, pointer2, pointer3)

	// Types of array
	var (
		array1 [2]int32
		array2 [2]string
		array3 [2]float32
	)

	fmt.Println(array1, array2, array3)

	// Types of slice
	var (
		slice1 []string
		slice2 []int
		slice3 []float32
	)

	fmt.Println(slice1, slice2, slice3)

	// Types of map
	var (
		map1 map[string]int
		map2 map[string]string
		map3 map[int]float32
		map4 map[float64]float32
		map5 map[float32]float32
		map6 map[string]float32
	)

	fmt.Println(map1, map2, map3, map4, map5, map6)

	// Types of interface
	var (
		interface1 interface{}
		interface2 personInterface
	)

	fmt.Println(interface1, interface2)

	// Types of struct
	var (
		struct1 person
	)

	fmt.Println(struct1)

}

type person struct {
	name string
	age  int
}

type personInterface interface {
	DisplayNumber()
}
