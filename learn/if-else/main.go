package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x > y {
		fmt.Println("x is greater than y")
	} else if x == y {
		fmt.Println("x is equal to y")
	} else {
		fmt.Println("x is less than y")
	}
}
