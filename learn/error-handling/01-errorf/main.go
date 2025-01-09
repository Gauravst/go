package main

import "fmt"

func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot cannot divide %d by zero", x)
	}

	return x / y, nil
}

func main() {

	var x int
	var y int

	fmt.Println("Enter first number :- ")
	fmt.Scanln(&x)

	fmt.Println("Enter second number :- ")
	fmt.Scanln(&y)

	ans, err := divide(x, y)
	if err != nil {
		fmt.Println("Error is :- ", err)
		return
	}

	fmt.Println("Answer is :- ", ans)
}
