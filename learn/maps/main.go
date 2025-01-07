package main

import "fmt"

func main() {

	personAge := map[string]int{
		"Gaurav": 30,
		"ram":    34,
	}

	fmt.Println("gaurav's age", personAge["Gaurav"])

}
