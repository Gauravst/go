package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, Go Routine!")
}

func main() {
	go sayHello()
	time.Sleep(time.Second)
}
