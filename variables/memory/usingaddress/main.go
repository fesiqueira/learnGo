package main

import (
	"fmt"
)

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")

	// Scan receives the memory address that the variable should be stored
	// It's like to say to Scan:
	// "Take it and put it inside the meters's box"
	fmt.Scan(&meters)
	// And now meters has a value

	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
}
