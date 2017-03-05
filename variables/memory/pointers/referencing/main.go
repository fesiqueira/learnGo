package main

import "fmt"

func main() {
	a := 10

	// Shows a's value
	fmt.Println(a)
	// Shows a's memory address
	fmt.Println(&a)

	// b is an pointer to the a's memory address and not to a's value
	// * is part of the type, so b is a *int (integer pointer)
	// It can store memory addresses for integer, and only for integers
	// b is a integer pointer
	var b *int = &a

	// When b is printed, it's referencing to the a's address.
	// So it's value is the same as the &a (a's memory address)
	fmt.Println(b)
}
