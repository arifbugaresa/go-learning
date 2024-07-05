package modules

import "fmt"

// LearningConditional
// Conditionals are a method used in computer programs to make decisions based on whether certain conditions are true or false.
func LearningConditional() {
	IfElse()
	SwitchCase()
}

func IfElse() {
	// conditional if
	if true {
		fmt.Println("jalankan code")
	}

	var mood = "happy"
	if mood == "happy" {
		fmt.Println("hari ini aku bahagia!")
	}

	// conditional if else
	var grocheries = "open"
	if grocheries == "open" {
		fmt.Println("saya akan membeli telur dan buah")
	} else {
		fmt.Println("minimarketnya tutup")
	}

	// temporary variable
	if status, minuteRemainingToOpen := "close", 5; status == "open" {
		fmt.Println("I will buy eggs and fruits.")
	} else if minuteRemainingToOpen <= 5 {
		fmt.Println("The minimarket will open soon, I'll wait.")
	} else {
		fmt.Println("The minimarket is closed, I'll go back home.")
	}
}

func SwitchCase() {
	var buttonPushed = 1
	switch buttonPushed {
	case 1:
		fmt.Println("Turn off the TV!")
	case 2:
		fmt.Println("Lower the TV volume!")
	case 3:
		fmt.Println("Increase the TV volume!")
	case 4:
		fmt.Println("Mute the TV!")
	default:
		fmt.Println("Nothing happened")
	}

	// fallthrough
	var point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
