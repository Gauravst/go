package main

import (
	"fmt"
	"time"
)

func printNumbers(label string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(label, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers("Goroutine 1")
	go printNumbers("Goroutine 2")

	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
}
