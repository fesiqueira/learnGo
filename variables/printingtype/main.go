package main

import "fmt"

func main() {

	// Variables initialization (declaration and value assignment at the same time)
	a := 10
	b := "Felipe"
	c := 3.14
	d := true

	// For more info about the verbs: check the "fmt" package documentation

	fmt.Printf("Printing types \n")
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
