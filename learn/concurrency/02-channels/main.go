package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "My message 01"
		messages <- "My message 02"
	}()

	msg1 := <-messages
	msg2 := <-messages

	fmt.Println("first message is :- ", msg1)
	fmt.Println("second message is :- ", msg2)

}
