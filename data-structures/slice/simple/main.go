package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}

	// Slices are lists ("dynamic arrays"), it means that the length of a slice can increse or decrease
	// The type printed is a slice of int, because there's no length inside the []
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)
}
