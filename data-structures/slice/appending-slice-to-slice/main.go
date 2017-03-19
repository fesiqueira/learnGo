package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	otherSlice := []int{6, 7, 8, 9}

	// mySlice is expecting a int, and not a []int.
	// So using the variadic notation ... will iterate the append over the otherSlice content
	// Passing every item inside the otherSlice individually
	mySlice = append(mySlice, otherSlice...)

	fmt.Println(mySlice)
}
