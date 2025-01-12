package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
}

// Define a struct for Circle
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Define a struct for Rectangle
type Rectangle struct {
	Width, Height float64
}

// Implement the Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	// Both Circle and Rectangle are treated as Shape
	printArea(c) // Output: Area: 78.5
	printArea(r) // Output: Area: 12
}
