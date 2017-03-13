package main

import "fmt"

func main() {
	// Defer appends the execution of whatever is on the line
	// that defer is called to an execution stack
	// So right before the function that contains a defer inside
	// or the program exit, the stack is executed,
	// following the LIFO (Last In First Out)
	defer fmt.Println(1)
	fmt.Println(2)
}
