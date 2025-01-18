package main

import (
	"fmt"
	"net/http"
)

func testHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello Gaurav")
}

func main() {
	fmt.Println("Server is running on http://localhost:8080")

	http.HandleFunc("/", testHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
