package main

import (
	"fmt"
)

// In Go we don't capitalize the constants like other languages!
// Initializing a constant, which is unchangeable
// The string type is redundant, once it's inferred when the constant is initialized
// The constants MUST be initialized (declared and have a value assigned)
const p string = "Felipe Siqueira"

func main() {
	// Constant must be initialized as below,
	// It MUST start with the const keyword
	const q = 42

	fmt.Println("p:", p)
	fmt.Println("q:", q)
}
