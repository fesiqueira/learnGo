package main

import "fmt"

func main() {
	// 2^8 = 256
	// So a byte is a uint8
	// And can store only from 0 to 255 (u stands for unsigned, so uint haven't nevatives)
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		// The for initializes the i as a int, but the array x is specting a uint8 (a byte)
		// So i is converted: byte(i). But it still a integer, but only with 8 bits
		// And the value stills a numeric value
		x[i] = byte(i)
	}
	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
}
