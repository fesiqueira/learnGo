package main

import "fmt"

// Function Wrapper returns a function and a integer
// The function within wrapper returns a integer
// The x initialized in the wrapper is on the outer scope, relative to the
// anonymous function that is assigned to increment
// So, when the increment() is called, it will increment the x (that is declared)
// on the outer scope, and is accessible for the increment()
// and it will return x

// As the inner function returns a integer, the outer function(the wrapper function),
// needs to return integer too

// The x within wrapper is initialized in the outer scope, relative to the
// anonymous inner function, that is assigned to increment
// So when increment() is called, it have access to the x variable

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	// When wrapper is called, it will return a anonymous function
	// that is assigned to increment, so now increment is a function
	// And increment now returns a integer, so wrapper must return a integer too
	// The x declared inside the wrapper is only accessible inside it's own curlies(block scope)
	// So it can only be accessed by the increment(that is the wrapper scope)
	increment := wrapper()
	fmt.Println("First call:", increment())
	fmt.Println("Second call:", increment())
}
