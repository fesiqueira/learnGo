package main

import (
	"fmt"
)

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println("First call:", increment())
	fmt.Println("Second call:", increment())
}
