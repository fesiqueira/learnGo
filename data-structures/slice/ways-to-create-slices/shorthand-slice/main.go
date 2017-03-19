package main

import "fmt"

func main() {
	// Every new item should be appended to the slice
	student := []string{}
	students := [][]string{}

	// student[0] = "Felipe" // Index out of range, because there's no underlying array

	fmt.Println(student)
	fmt.Println(students)

	// In this case the slices were initialized, both of them have length and capacity 0
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
