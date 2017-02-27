package main

import "fmt"

//Variable initialized on the package scope, accessible from any file that is part of the same package
var x int = 42

func main() {
	fmt.Println(x)
	printX()

	// Initializing y variable. The line 13 can't see the y, because it's declared below the print.
	fmt.Println(y)
	y := "Felipe"
}

func printX() {
	fmt.Println(x)

	// The line below can't print y because y is declared in the func main scope (inside curly braces {} )
	fmt.Println(y)
}
