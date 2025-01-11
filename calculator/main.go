package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func mul(x int, y int) int {
	return x * y
}

func dev(x int, y int) int {
	return x / y
}

func main() {
	var x int
	var y int
	var operator int

	fmt.Println("Calculator...........")
	fmt.Println("Enter first number :----")
	fmt.Scanln(&x)

	fmt.Println("Enter second number :------")
	fmt.Scanln(&y)

	fmt.Println("Enter Your Task :--")
	fmt.Println("1) +")
	fmt.Println("2) -")
	fmt.Println("3) *")
	fmt.Println("4) /")
	fmt.Scanln(&operator)

	switch operator {
	case 1:
		fmt.Println("Your ans is :- ", add(x, y))

	case 2:
		fmt.Println("Your ans is :- ", sub(x, y))

	case 3:
		fmt.Println("Your ans is :- ", mul(x, y))

	case 4:
		fmt.Println("Your ans is :- ", dev(x, y))

	default:
		return
	}

}
