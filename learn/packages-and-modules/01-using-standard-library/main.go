package main

import (
	"fmt"
	"math"

	"../02-creating-custom-packages/check"
)

func main() {
	fmt.Println(math.Sqrt(4))
	fmt.Println(check.Hello("Go"))
}
