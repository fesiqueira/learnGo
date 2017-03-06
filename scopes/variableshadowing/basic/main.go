package main

import "fmt"

// When a function is declared, it can have or not PARAMETERS
// And when a function is called, it can receive or not ARGUMENTS

// Defined in the package scope, not accessible globally nor exported (named with lower case).
// Accessible from any file that is part of the package
func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)

	// max is now the result of the return from the max function. So max is not a function anymore
	fmt.Println(max)

	// max is not a function anymore
	a := max(7)
}
