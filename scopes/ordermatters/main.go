package main

import (
	"fmt"
)

func main() {
	// x scope is limited at the line it's declared and below. So it's inaccessible above it's declaration. Order matters
	fmt.Println(x)
	x := 10
	fmt.Println(y)
}

// This variable is initialized in the package scope, so it's accessible inside any file of the package,
// declaration order doesn't matter here

// The package scope (outer scope) is enclosuring the block scope (inner scope), like russian dolls. It's closure!
var y = 42
