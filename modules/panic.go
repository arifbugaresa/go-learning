package modules

import "fmt"

func EndApp() {
	fmt.Println("application is complete")
}

func RunApp(isError bool) {
	if isError {
		panic("Something went wrong")
	}
}

func LearningPanic() {
	isError := false

	RunApp(isError)

	EndApp()
}
