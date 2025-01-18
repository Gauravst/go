package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./file-handling/05-renaming-file/"

	err := os.Rename(path+"temp.txt", path+"example.txt")
	if err != nil {
		fmt.Println("Error :- ", err)
		return
	}

}
