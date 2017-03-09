package main

import "fmt"

func main() {
	x := 0
	switch x {
	case 1, 0:
		fmt.Println("The case above works as an OR")
	case 2:
		fmt.Println("The second doens't run")
	default:
		fmt.Println("None")
	}
}
