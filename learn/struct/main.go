package main

import "fmt"

func main() {
	type Person struct {
		name string
		age  int
		id   string
	}

	p1 := Person{
		name: "gaurav",
		age:  34,
		id:   "fners",
	}

	p2 := Person{
		name: "ram",
		age:  354,
		id:   "fg",
	}

	fmt.Println(p1)
	fmt.Println(p2)

}
