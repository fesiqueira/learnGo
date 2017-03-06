package main

import "fmt"

func main() {
	// Initializing a function with a character between simple quotations marks
	// it's type will be an int32 (a rune), the decimal representation from the ASCII table
	a := 'i' // 105

	fmt.Printf("%T \n", a) // a is a int32 (a rune)
	fmt.Printf("%d \n", a)

	// Converting to a rune is redundant, int32 is a rune
	fmt.Printf("%T \n", rune(a))
}
