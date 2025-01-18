package main

import (
	"fmt"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	fmt.Println("Sever running on port 8080")

	http.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error - ", err)
	}
}
