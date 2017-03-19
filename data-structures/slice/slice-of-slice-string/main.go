package main

import "fmt"

func main() {
	// Records is a slice of slice of string
	// Records is a slice that stores slices
	records := make([][]string, 0)
	student1 := make([]string, 4)

	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	// The records appending don't need the variadic notation because it's expecting a slice of strings
	records = append(records, student1)

	student2 := make([]string, 4)

	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"

	records = append(records, student2)

	fmt.Println(records)
}
