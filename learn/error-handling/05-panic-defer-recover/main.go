package main

import "fmt"

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("Something went worng")
}

func main() {
	fmt.Println("test 1")
	test()
	fmt.Println("test 2")
}
