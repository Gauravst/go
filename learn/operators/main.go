package main

import "fmt"

func main() {
	x := 20
	y := 2

	// additions
	fmt.Println("additions", x+y)

	// subraction
	fmt.Println("subraction", x-y)

	// multiple
	fmt.Println("multiple", x*y)

	// division
	fmt.Println("division", x/y)

	// modules
	fmt.Println("modules", x%y)

	// increment
	x++
	fmt.Println("increment", x)

	// decrement
	y--
	fmt.Println("decrement", x)

}
