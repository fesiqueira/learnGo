package main

import "fmt"

// That is another way to pass multiple values to a variadic function
// Making the function signature receive only slices/arrays of the type
// And passing the arguments without the three dots (data...)

func main() {
	data := []float64{1, 2, 3, 4, 5}

	// The multiple arguments are being passed directly, without the three dots
	n := average(data)
	fmt.Println(n)
}

// The function receives only a slice/array of type float64
func average(n []float64) float64 {
	var t float64
	for _, v := range n {
		t += v
	}
	return t / float64(len(n))
}
