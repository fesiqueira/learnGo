package main

import "fmt"

func main() {
	n := average(10, 11, 12, 14, 15, 16)
	fmt.Println(n)
}

func average(numbers ...float64) float64 {
	total := 0.0
	for _, v := range numbers {
		total += v
	}
	return total / float64(len(numbers))
}
