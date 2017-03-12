package main

import "fmt"

// The following function returns another function
// So, when this function is invoked, it must be assigned to a variable
// And the variable will turn into a function.

// The outer function must return the same types of the inner function
// In this example, the anonymous function returns a string, so the makeGreeter must return a string too
func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())

	// It's typed as a func() string
	// It's a function that returns a string
	fmt.Printf("%T\n", greet)

	// It's a function that returns a function that return a string
	fmt.Printf("%T\n", makeGreeter)
}
