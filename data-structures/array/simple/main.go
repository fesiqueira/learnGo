package main

import "fmt"

func main() {
	// If there's no number inside the brackets, it's a slice. If there's numbers inside
	// it's an array

	// Arrays are static data-structures, they length can't be changed.
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
