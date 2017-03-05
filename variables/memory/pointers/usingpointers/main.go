package main

import "fmt"

func main() {
	a := 10

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Println(*b)

	// It means: "Assign 11 to the value of this memory address"
	// Because b is a pointer to a
	// Any change made on a will affect b
	// And any change made on b will affect a
	*b = 11

	// Now a is equals 11
	fmt.Println(a)
}
