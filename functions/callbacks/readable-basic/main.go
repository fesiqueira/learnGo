package main

import "fmt"

func visit(callback func(n int), numbers ...int) {
	for _, numbers := range numbers {
		callback(numbers)
	}
}

// This file does exactly the same thing that the basic/ does
// But I think this one is more "readable"
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	callback := func(n int) {
		fmt.Println(n)
	}

	visit(callback, numbers...)
}
