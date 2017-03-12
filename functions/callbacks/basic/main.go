package main

import "fmt"

// visit receives a slice of integers and a function that receives a integer as parameters
func visit(numbers []int, callback func(int)) {

	// invokes callback() for each number inside the []int
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	// The callback() has it's "value assigned" just when the visit() is invoked
	// In this case, callback() will be equivalent as the following:
	//
	// func callback(n int) {
	// fmt.Println(n)
	// }

	// Callback is passing a function as an argument
	visit([]int{1, 2, 3, 4, 5, 6}, func(n int) {
		fmt.Println(n)
	})
}
