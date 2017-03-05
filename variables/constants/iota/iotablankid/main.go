package main

import "fmt"

// Assigning the first iota to a blank identifier, it will throw the current value away
// But the following iotas will remember the value, and will keep incrementing
const (
	_ = iota
	b = iota
	c = iota
)

func main() {
	// fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
