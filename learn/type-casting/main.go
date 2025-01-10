package main

import (
	"fmt"
	"strconv"
)

func main() {

	// int to float
	var a int = 5
	var b float32 = float32(a)

	fmt.Printf("int to float %d to %f \n", a, b)

	// float to int
	var c float64 = 5.55
	var d int64 = int64(c)

	fmt.Printf("float to int %f to %d \n", c, d)

	// string to int
	var e string = "123"
	ans, err := strconv.Atoi(e)

	if err != nil {
		fmt.Println("Conversion error:", err)
	} else {
		fmt.Println("string to int - ", ans)
	}

	// int to string
	var f int = 233
	g := strconv.Itoa(f)
	fmt.Printf("int %d to string %q", f, g)

}
