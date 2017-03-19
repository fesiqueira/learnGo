package main

import "fmt"

func main() {
	var x [256]int

	// The type is an array of int, because there's the length inside the []
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
