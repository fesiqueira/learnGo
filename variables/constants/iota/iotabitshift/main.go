package main

import "fmt"

// The << represents bit shifting
// It's like multiplying a number by the 2 raised to power of the number you shifted

// 1 << 10
// Is the same to raise 2 to the 10 and then multiply the result (1024) to the 1
// 1 << 10 is equals to 1 * (2^10)
// That is equals to 1 * 1024
// So 1 << 10 = 1024

// 1 in binary is 1
// When 1 << 10, it will drag the 1 ten bits to the left
// Resulting: 10000000000
// That number in binary is equals to 1024

const (
	_  = iota
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
)

func main() {
	fmt.Println("Binary\t\tDecimal")
	fmt.Printf("%b\t", kb)
	fmt.Printf("%d\n", kb)
	fmt.Printf("%b\t", mb)
	fmt.Printf("%d\n", mb)
}
