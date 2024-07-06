package modules

import "fmt"

func logging() {
	fmt.Println("logging")
}

func LearningDefer() {
	defer logging()

	fmt.Println("Run App")
}
