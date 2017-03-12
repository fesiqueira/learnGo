package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// The example below is an for each.
	// The range keyword gets the Index and the Value of a slice/array and sets the index to the first variable
	// And sets the values to the second variable

	// If you want only the value, you could replace Index with a blank identifier
	// If yoy want only the Index, you could remove the value variable, and attribute the range directly: index := range x {}
	for index, value := range x {
		fmt.Println("Index:", index, "\t", "Value:", value)
	}
}
