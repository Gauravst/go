package main

import (
	"fmt"
	"sync"
	"time"
)

func bakeCookies(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this friend as done when finished

	fmt.Printf("Friend %d is baking cookies\n", id)
	time.Sleep(time.Second) // Simulate time taken to bake cookies
	fmt.Printf("Friend %d is done baking cookies\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Mark a friend as "started"
		go bakeCookies(i, &wg)
	}

	wg.Wait() // Wait for all friends to finish
	fmt.Println("All friends are done baking!")
}
