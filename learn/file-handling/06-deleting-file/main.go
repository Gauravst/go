package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./file-handling/06-deleting-file/"

	err := os.Remove(path + "temp.txt")
	if err != nil {
		fmt.Println("Error deleting-file :-", err)
		return
	}
}
