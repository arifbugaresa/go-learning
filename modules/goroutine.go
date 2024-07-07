package modules

import (
	"fmt"
	"sync"
	"time"
)

func LearningGoroutines() {
	Goroutines()
	Channel()
	ClosingChannel()
	ChannelBuffer()
	WaitGroup()
}

func Goroutines() {
	go say("Hello, world!")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Channel() {
	numberChannel := make(chan int)

	go Display(numberChannel)

	numberInteger := <-numberChannel

	fmt.Println("integer channel value :", numberInteger)
}

func Display(ch chan int) {
	fmt.Println("this is from a goroutine")
	ch <- 10
}

func ClosingChannel() {
	ch := make(chan int)

	go cetak(ch)

	for {
		data, ok := <-ch
		if ok == false {
			break
		}
		fmt.Printf("Data di terima %v\n", data)
	}
}

func cetak(ch chan int) {
	for index := 0; index < 10; index++ {
		ch <- index
	}
	close(ch)
}

func ChannelBuffer() {
	ch := make(chan int, 3)

	ch <- 6
	ch <- 7
	ch <- 5

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func WaitGroup() {
	GoroutineWithoutWaitGroup()
	GoroutineWithWaitGroup()
}

func GoroutineWithoutWaitGroup() {
	go printText("Halo")
	go printText("Dunia")

	time.Sleep(500 * time.Millisecond)
}

func printText(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
}

func printTextWithWaitGroup(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}

	wg.Done()
}

func GoroutineWithWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printTextWithWaitGroup("Halo", &wg)

	wg.Add(1)
	go printTextWithWaitGroup("Dunia", &wg)

	wg.Wait()
}
