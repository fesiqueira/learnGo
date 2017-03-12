package main

import "fmt"

func main() {

	// Variadic functions can receive an slice/array of elements as argument over a bunch of arguments separated by commas
	data := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// n will receive the return value from average function, which is a Variadic function
	// But the arguments to that function are being passed with the following notation: sliceName...

	// The "data..." notation means that the arguments that is being passed is already a slice/array, even if the
	// types diverge (the arguments are a slice of float64, and not a simple float64 value)
	// Each value within the slice will be passed one by time, until all the values be passed to the variadic function
	n := average(data...)
	fmt.Println(n)
}

// The function signature changes nothing to receive a slice/array over the comma-separated values
func average(numbers ...float64) float64 {
	var t float64
	for _, v := range numbers {
		t += v
	}
	return t / float64(len(numbers))
}
