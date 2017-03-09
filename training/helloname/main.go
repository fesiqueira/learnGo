package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
