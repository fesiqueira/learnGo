package main

import "fmt"

func wrapper() func() int {
	// x is only accessible on the inner scopes
	// the anonymous function will Print the x because it's accessible on the inner scopes
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}
