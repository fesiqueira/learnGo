package main

import "fmt"

// zero() receives a memory address, in this case: &x
func zero(x *int) {
	// *x direferences the memory address, and gets the x's value
	// The value that is stored in the x's memory address
	// *x = 0 changes the value store in the x's address
	// So when x is printed in the func main, it value will be 0
	*x = 0
}

func main() {
	x := 5

	// Passing x's memory address to zero()
	zero(&x)

	// x's value were changed by the *x
	// Changed the value stored in the x's memory address
	fmt.Println(x)
}
