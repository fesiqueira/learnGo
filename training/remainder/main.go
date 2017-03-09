package main

import (
	"fmt"
)

func main() {
	var large, small int
	fmt.Println("Enter a large number and a small number: ")
	fmt.Scan(&large)
	fmt.Scan(&small)
	fmt.Println(large % small)
}
