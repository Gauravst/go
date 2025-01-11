package main

import "fmt"

func main() {
	numbers := make(chan int, 4)

	numbers <- 1
	numbers <- 2
	numbers <- 3
	numbers <- 4

	fmt.Printf("\n1st element of number is :- %d", <-numbers)
	fmt.Printf("\n2nd element of number is :- %d", <-numbers)
	fmt.Printf("\n3rd element of number is :- %d", <-numbers)
	fmt.Printf("\n4th element of number is :- %d", <-numbers)
}
