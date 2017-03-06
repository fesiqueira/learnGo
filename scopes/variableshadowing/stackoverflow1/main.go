package main

import "fmt"

func main() {
	x := 3

	double := func() {
		// The x below is the same x initialized in the func main() scope
		fmt.Println(&x) // x's memory address
		x = x * 2
	}
	double()
	fmt.Println(&x) // x's memory address
	fmt.Println(x)
}
