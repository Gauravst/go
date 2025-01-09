package main

import "fmt"

// this is type of err
type errorType struct {
	Dividend int
	Divisor  int
	Message  string
}

// this is function of error to create custom error
func (e *errorType) Error() string {
	return fmt.Sprintf("Cannot divide %d by %d : %s", e.Dividend, e.Divisor, e.Message)
}

// this is function for divide
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, &errorType{
			Dividend: x,
			Divisor:  y,
			Message:  "division by zero",
		}
	}

	return x / y, nil
}

// this is main function
func main() {

	var x int
	var y int

	fmt.Println("Enter first number :- ")
	fmt.Scanln(&x)

	fmt.Println("Enter second number :- ")
	fmt.Scanln(&y)

	ans, err := divide(x, y)

	if err != nil {
		if de, ok := err.(*errorType); ok {
			fmt.Println("custom error caught :- ", de)
		} else {
			fmt.Println("Error :- ", err)
		}
		return
	}

	fmt.Println("Answer is :- ", ans)

}
