package main

import "fmt"

// x is initialized on the package scope, it's accessible from any file of the same package, and accessible from the inner scopes
var x int

func increment() int {
	// As x is accessible from here, because it was initialized at the package level, the increment
	// increments the same x that was initialized right above this function signature

	// If the return was a named return, like (x int), x would be re-declared, so the increment would return another x
	// And if the x was printed, it would show 0, because it was unchaged by increment()
	// The fmt.Println(increment()) would print 1, because x was re-declared in the increment() scope, and the compiler
	// initialized it's value with the zero value for integers, that is 0
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(x)
}
