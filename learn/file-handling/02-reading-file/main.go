package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("./file-handling/02-reading-file/test.txt")
	if err != nil {
		fmt.Println("error reading-file :- ", err)
		return
	}

	fmt.Println(string(content))

}
