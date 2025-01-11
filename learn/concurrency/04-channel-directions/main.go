package main

import "fmt"

func sendData(ch chan<- string) {
	ch <- "Hello from sendData function"
}

func receiveData(ch <-chan string) {
	msg := <-ch
	fmt.Println(msg)
}

func main() {
	ch := make(chan string)

	go sendData(ch)
	receiveData(ch)
}
