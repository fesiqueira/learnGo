package main

import (
	"fmt"
)

func main() {
	// x is initialized in the outer scope
	// It stills a block scope (because it's within curly braces), but it's an outer scope
	// relative to the scope below, that is delimited by curly braces too.
	x := 10
	fmt.Println("Outer scope:", x)

	// Delimiting an inner scope
	// Any function/variable declared inside the curlies will only be accessible inside these curlies
	// Inner scope
	{
		fmt.Println("I'm accessible:", x)
		// y := 11
		// y is accessible, but only in it's own scope (every line below it's initialization
		// until the next closing curly brace}
		y := 11
		fmt.Println("I'm accessible here:", y)
	}
	// fmt.Println("I'm not accessible:", y)
}
