package main

import (
	"fmt"
)

func main() {

	// The following variable is initialized with an anonymous function
	// So, that is no function identifier

	// That's the only way to have a function inside another function
	// It's called a Function Expression
	greeting := func() {
		fmt.Println("Hello World")
	}
	greeting()
	fmt.Printf("%T\n", greeting)
}
