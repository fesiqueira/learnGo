package main

import (
	"fmt"
)

// Variable declaration: b as type String
var g string

// Variable declaration and value assignment: variable initialization
var e = 20

func main() {
	// Variable value assignment
	g = "Siqueira"

	// Variables declaration and assignment simultaneously: variable initialization
	a := 1
	b := "Felipe"
	c := 3.14
	d := true

	// For more info about the verbs: check the "fmt" package documentation

	fmt.Printf("Printing values: \n")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
