package main

import "fmt"

func main() {

	var slice1 []int = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	slice2 := []int{5, 6}
	fmt.Println(slice2)

	fmt.Println("first element of slice2 -:- ", slice2[0])

	fmt.Println(append(slice2, 7, 8))

	fmt.Println(append(slice1, slice2...))
	slice1[0] = 99
	fmt.Println(slice1)

}
