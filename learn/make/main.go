package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
