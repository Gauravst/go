package main

import (
	"fmt"
	"sync"
)

var (
	cookies = 0
	mutex   sync.Mutex
)

func addCookie() {
	mutex.Lock()
	cookies += 1
	mutex.Unlock()
}

func main() {
	addCookie()
	fmt.Println("Cookies:", cookies)
}
