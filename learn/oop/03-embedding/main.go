package main

import "fmt"

// Define a struct to be embedded
type Engine struct {
	Horsepower int
}

// Define another struct that embeds Engine
type Car struct {
	Engine // Embedded struct
	Brand  string
}

func main() {
	myCar := Car{
		Engine: Engine{Horsepower: 150},
		Brand:  "Toyota",
	}

	// Access fields of the embedded struct directly
	fmt.Println("Brand:", myCar.Brand)           // Output: Brand: Toyota
	fmt.Println("Horsepower:", myCar.Horsepower) // Output: Horsepower: 150
}
