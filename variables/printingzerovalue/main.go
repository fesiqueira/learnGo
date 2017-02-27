package main

import "fmt"

func main() {
	// Variable declaration (without assigning a value)
	var a string
	var b int
	var c float64
	var d bool

	fmt.Printf("Printing the zero-values \n")
	fmt.Printf("%T: %v \n", a, a)
	fmt.Printf("%T: %v \n", b, b)
	fmt.Printf("%T: %v \n", c, c)
	fmt.Printf("%T: %v \n", d, d)
}
