package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)

	// x stills 5 because zero "created a copy" of x and assigned the 0
	// to the "x copy"
	// The original x is unchanged
	fmt.Println(x)
}
