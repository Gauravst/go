package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code %d : %s", e.Code, e.Message)
}

func someFunc() error {
	return &MyError{
		Code:    500,
		Message: "Something went worng",
	}
}

func main() {
	err := someFunc()

	var myErr *MyError
	if errors.As(err, &myErr) {
		fmt.Printf("Custom error: Code=%d, Message=%s\n", myErr.Code, myErr.Message)
		return
	}

	fmt.Println("Some other error occurred.")
}
