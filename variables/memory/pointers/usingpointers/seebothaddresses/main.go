package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x)
	x = 0
}

// See that the x's address are different in both scopes
// It means that the both x variables are not the same variable
// That's why x stills 1 after calling zero(x)
func main() {
	x := 1
	fmt.Printf("%p\n", &x)
	zero(x)
	fmt.Println("X:", x)
}
