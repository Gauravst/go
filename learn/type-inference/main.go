package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int = 5
	var y string = "Gaurav"

	fmt.Println("x type is - ", reflect.TypeOf(x))
	fmt.Printf("y type is - %T", y)

}
