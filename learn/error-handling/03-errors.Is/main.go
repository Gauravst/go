package main

import (
	"errors"
	"fmt"
)

var someError = errors.New("not found")

func someFunction() error {
	return fmt.Errorf("failed to find item : %w", someError)
}

func main() {

	err := someFunction()

	if errors.Is(err, someError) {
		fmt.Println("item was not found")
		return
	}

	fmt.Println("some other error occurred")

}
