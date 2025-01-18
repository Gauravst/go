package main

import (
	"fmt"
	"net/http"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to home page")
}

func aboutPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Welcome to about page")
}

func main() {
	fmt.Println("Sever is running on port 8080")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server :- ", err)
	}

}
