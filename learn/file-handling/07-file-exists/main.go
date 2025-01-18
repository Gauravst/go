package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./file-handling/07-file-exists/"

	_, err := os.Stat(path + "temp.txt")

	if os.IsNotExist(err) {
		fmt.Println("file does not exist")
		return
	}

	fmt.Println("file exist")

}
