package modules

import (
	"errors"
	"fmt"
	"strconv"
)

func LearningError() {
	BasicError()
	CustomError()
}

func BasicError() {
	var input = "aa"

	// convert string to number
	_, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CustomError() {
	var input = "aa"

	// convert string to number
	_, err := strconv.Atoi(input)
	if err != nil {
		err = errors.New(input + " is not number")
		fmt.Println(err.Error())
	}
}
