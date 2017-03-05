package main

import "fmt"

func main() {
	x := 10

	fmt.Println("x:", x)
	// & represents that memory address where x is stored
	fmt.Println("x's address:", &x)
}
