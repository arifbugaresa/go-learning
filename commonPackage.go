package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LearningCommonGoPackages
// Various types of packages commonly used
func LearningCommonGoPackages() {
	packageString()
	packageStrconv()
}

// packageString
// the strings package is used to modify the value of a string variable.
func packageString() {
	// strings index
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	// strings replace
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	// strings repeat
	var str1 = strings.Repeat("na", 4)
	fmt.Println(str1) // "nananana"

	// strings toLower
	var str2 = strings.ToLower("aLAy")
	fmt.Println(str2) // "alay"

	// strings toUpper
	var str3 = strings.ToUpper("eat!")
	fmt.Println(str3) // "EAT!"
}

// packageStrconv
// used for converting between data types
func packageStrconv() {
	// convert string to int
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num) // 124
	}

	// convert int to string
	var num1 = 124
	str1 := strconv.Itoa(num1)
	fmt.Println(str1) // "124"

	// convert string to float
	var str2 = "24.12"
	num2, _ := strconv.ParseFloat(str2, 32)

	fmt.Println(num2) // 24.1200008392334

	// convert string to boolean
	var str3 = "true"
	var bul3, _ = strconv.ParseBool(str3)

	fmt.Println(bul3)
}
