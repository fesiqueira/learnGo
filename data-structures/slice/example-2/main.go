package main

import "fmt"

func main() {
	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)

	// The slicing below will print a slice from the 2 index, until the 3 index (excluding the 4)
	fmt.Println(mySlice[2:4]) // Slicing a slice

	// A default kind of access, really similar to other languages
	fmt.Println(mySlice[2]) // index access; accessing by index

	// It's possible to do a slicing on a string because a string is a slice of bytes
	// A string is made up of runes
	// A rune is a unicode code point
	// A unicode code point is 1 to 4 bytes

	// String are made up of runes, and runes are made up of bytes.
	// So strings are made up of bytes. A string is a bunch of bytes; A slice of bytes

	// This will print the byte at the index 2, which is 83, or 'S' (the decimal representation for the S, on the ASCII table)
	fmt.Printf("%T\n", "mySlice"[2]) // index access; acessing by index
	fmt.Println("mySlice"[2])        // index access; acessing by index
}
