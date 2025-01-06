package main

import "fmt"

func main() {
	x := 10

	switch x {
	case 1:
		fmt.Println("x is 1")

	case 5:
		fmt.Println("x is 5")

	default:
		fmt.Println("x is unknown")
	}

}
