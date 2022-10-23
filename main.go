package main

import (
	"fmt"
	"time"
)

func main() {
	index := make(chan int)

	go MyCounter(0, index)

	for i := range index {
		fmt.Printf("Counter: %d\n", i)
	}
}

func MyCounter(start int, index chan int) {
	var lastIndex int
	defer func() {
		// you must recover from panic inside a go routine not outside of it
		if err := recover(); err != nil {
			fmt.Println("Recovered in MyCounter", err)
			go MyCounter(lastIndex+2, index)
		}
	}()

	for i := start; i < 10; i++ {
		time.Sleep(250 * time.Millisecond)

		if i == 5 {
			panic("Yikes!")
		}

		lastIndex = i
		index <- i
	}
	close(index)
}
