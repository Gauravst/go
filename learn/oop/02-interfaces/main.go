package main

import "fmt"

// Define an interface
type Vehicle interface {
	Drive()
}

// Car type implements Vehicle interface
type Car struct {
	Brand string
}

func (c Car) Drive() {
	fmt.Println(c.Brand, "is driving")
}

// Bike type implements Vehicle interface
type Bike struct {
	Model string
}

func (b Bike) Drive() {
	fmt.Println(b.Model, "is driving")
}

func main() {
	var v Vehicle

	v = Car{Brand: "Toyota"}
	v.Drive() // Output: Toyota is driving

	v = Bike{Model: "Yamaha"}
	v.Drive() // Output: Yamaha is driving
}
