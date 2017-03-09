package main

import (
	"fmt"
)

func main() {
	x := 0
	// Go automatically sets the break statement after the case *:
	// When the fallthough is needed, it must be specified
	switch x {
	case -1:
		fmt.Println("First option")
	case 0:
		fmt.Println("Second option")
	case 1:
		fmt.Println("Third option")
	default:
		fmt.Println("Default option")
	}
}
