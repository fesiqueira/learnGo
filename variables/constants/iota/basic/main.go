package main

import (
	"fmt"
)

// When a iota is declared, it starts with 0,
// and each new iota will have the previous value + 1 (0, 1, 2, ..., n iotas)
// iotas can have it's value modified by mathematical operations before assignment to the const
const (
	a = iota
	b = iota
	c = iota
	g = iota
)

// If iotas are initialized in another "const scope" it will restart from 0
// So the following iotas start from 0 and goes to 2
const (
	d = iota
	e = iota
	f = iota
)

func main() {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("f:", f)

}
