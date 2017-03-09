package main

import "fmt"

func main() {
	x := 0
	switch x {
	// fallthrough makes the next case be runned right after the case that
	// constains the keyword fallthrough
	case 0:
		fmt.Println("First")
		fallthrough
	case 1:
		fmt.Println("Second is printed too due to fallthrough")
	case 2:
		fmt.Println("Third")
	default:
		fmt.Println("None")
	}
}
