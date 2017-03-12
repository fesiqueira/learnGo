package main

import "fmt"

func filter(numbers []int, callback func(n int) bool) []int {
	var filtered []int
	for _, n := range numbers {
		// callback returns true or false, depending on the internal comparison, on the line 23
		if callback(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// callback returns true or false, depending on the comparison: if n is greater than 1, true
	// if it returns false, the line 9 isn't executed, and the numbers is not appended to the filtered slice
	callback := func(n int) bool {
		return n > 1
	}
	filtered := filter(numbers, callback)
	fmt.Println(filtered)
}
