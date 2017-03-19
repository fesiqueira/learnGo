package main

import "fmt"

func main() {
	// Every new item should be appended to the slice
	var student []string
	var students [][]string

	// student[0] = "Felipe" // Index out of range, because there's no underlying array

	fmt.Println(student)
	fmt.Println(students)

	// The slices are just declared, not initialized, so they are a nil (unitilialized)
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
