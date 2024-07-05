package modules

import "fmt"

// LearningLoop
// Looping is the process of repeatedly executing a block of code
func LearningLoop() {
	ForLoop1()
	ForLoop2()
	ForLoop3()
	ForLoop4()
	ForLoop5()
}

func ForLoop1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}
}

func ForLoop2() {
	var i = 0

	for i < 5 {
		fmt.Println("Angka", i)
		i++
	}
}

func ForLoop3() {
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 5 {
			break
		}
	}
}

func ForLoop4() {
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}

func ForLoop5() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}
