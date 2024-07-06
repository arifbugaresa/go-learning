package modules

import "fmt"

func EndAppWithRecover() {
	fmt.Println("application is complete")

	message := recover()
	fmt.Println("recovered from", message)
}

func RunAppWithRecover(isError bool) {
	if isError {
		panic("Something went wrong")
	}
}

func LearningRecover() {
	isError := false

	defer EndAppWithRecover()
	RunAppWithRecover(isError)
}
