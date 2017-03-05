package main

import (
	"fmt"
)

func main() {
	a := 10

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	// As b stores the a's memory address, the a's value can by accessed by
	// dereferencing b
	// It can be done using the * before the variable: *b
	fmt.Println(b)

	// *b will print the a's value, so if the a's value change, the *b will change too
	// Using thhe * is called dereferencing
	fmt.Println(*b)
}
