package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// type of item
type ItemType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// a list of item with item types
var Items = []ItemType{
	{ID: 1, Name: "Gaurav"},
	{ID: 2, Name: "Ram"},
}

func getItems(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(res).Encode(Items)
	if err != nil {
		http.Error(res, "Failed to encode items", http.StatusInternalServerError)
	}
}

func createNewItem(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var newItem ItemType
	err := json.NewDecoder(req.Body).Decode(&newItem)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	Items = append(Items, newItem)
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "Item created!")
}

func main() {
	fmt.Println("Server running on port 8080")

	http.HandleFunc("/items", getItems)
	http.HandleFunc("/items/create", createNewItem)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error on starting server :- ", err)
	}

}
