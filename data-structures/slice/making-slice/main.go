package main

import "fmt"

func main() {
	// Make will initialize a slice, a map or a channel

	// Making a slice: make([]sliceType, initialLength, capacityOfTheUnderlyingArray)
	mySlice := make([]int, 0, 5)

	// Make will initialize this slice with a starting length of 0, but a maximum size of 5
	// But when the slice size surpass the maximum size, "the computer" will initialized another underlying array
	// with the double of the previous size, in this case, 10
	// Then all the content of the previous array will be transfered to the new array of 10 indexes
	// And the slice will now point to the new array of 10.
	// If the slice surpass the new array size, a new array with the double of the size will be initialized again
	// And every step is repeated.

	// PS.: This have a performance cost

	fmt.Println("-------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}
