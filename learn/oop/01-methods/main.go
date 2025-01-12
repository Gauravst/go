package main

import "fmt"

type Car struct {
	Brand string
}

// method
func (c Car) Drive() {
	fmt.Println(c.Brand, "is driving")
}

func main() {
	myCar := Car{Brand: "Toyota"}
	myCar.Drive()
}
