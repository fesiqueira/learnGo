package main

import "fmt"

func main() {
	// x is initialized at the block scope of func main(), so it's accessible from any inner scope
	x := 10
	fmt.Println(x)
	{
		fmt.Println(x)

		// y is initialized at the block scope, inside the func main() scope
		// So it's accessible from any inner scope, but it's undefined from the outer scopes
		y := 0
		fmt.Println(y)
	}

	// y is undefined at this scope
	// fmt.Println(y)
}
