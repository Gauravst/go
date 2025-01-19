package main

type Person struct {
	name string
}

func main() {
	// Create a new Person object
	p1 := &Person{name: "Alice"}

	// p1 is in use here

	// Now, p1 goes out of scope
	p1 = nil

	// At this point, the Person object that p1 was pointing to is considered garbage
	// The garbage collector will eventually free that memory
}
