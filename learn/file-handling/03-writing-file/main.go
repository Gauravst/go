package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	basePath := "./file-handling/03-writing-file/"

	file, err := os.Create(basePath + "test-1.txt")
	if err != nil {
		fmt.Println("Error creating file :- ", err)
		return
	}

	writingContent := "Hello Go Sir"

	_, err = file.WriteString(writingContent)
	if err != nil {
		fmt.Println("Error writing file :- ", err)
		return
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking to beginning of file :- ", err)
		return
	}

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file :- ", err)
		return
	}

	fmt.Println(string(content))

}
