package main

import (
	"fmt"
	"os"
)

func main() {
	content := []byte("Hello Go")
	path := "./file-handling/04-writing-file/"

	err := os.WriteFile(path+"temp.txt", content, 0644)
	if err != nil {
		fmt.Println("Error writing to file", err)
		return
	}

	fmt.Println("File written successfully")
}
