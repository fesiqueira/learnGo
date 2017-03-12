package main

import "fmt"

func main() {
	variadic(10, 11, 12, 13, 14, 15)
}

// Variadic functions has 0 or n parameters
// The arguments passed to a variadic function are appended to a slice of the same type
func variadic(numbers ...int) {
	fmt.Printf("%T\n", numbers)
	fmt.Println(numbers)
}
