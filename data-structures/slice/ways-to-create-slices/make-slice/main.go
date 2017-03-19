package main

import "fmt"

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)

	// The student slice is 35 indexes long, so appending another item will make the length 36
	student = append(student, "Felipe")
	fmt.Println(student[35])
	fmt.Println(len(student))

	// In this case, the slices were initialized, and have a length and capacity of 35
	fmt.Println(student == nil)
	fmt.Println(students == nil)
}
