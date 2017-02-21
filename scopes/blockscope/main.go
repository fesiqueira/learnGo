package main

import "fmt"

func main() {
	// x is declared in a block scope (in this case, the func main scope)
	x := "Felipe"
	fmt.Println(x)
	printX()
}

func printX() {
	// It doesn't compile because x is declared inside a function (curly braces), that means: it's scope is limited by the braces
	fmt.Println(x)
}
