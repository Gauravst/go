package main

import (
	"errors"
	"fmt"
)

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return x / y, nil

}

func main() {
	var x int
	var y int

	fmt.Println("Enter first number to divide :- ")
	fmt.Scanln(&x)

	fmt.Println("Enter second number to divide :- ")
	fmt.Scanln(&y)

	ans, err := divide(x, y)
	if err != nil {
		fmt.Println("Error :- ", err)
	} else {
		fmt.Println("Answer :- ", ans)
	}
}
