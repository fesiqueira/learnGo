package main

import "fmt"

func main() {
	// Basically a rune is a character (a number)
	// A rune is an integer representation of an Unicode character
	// Rune is an alias to int
	for i := 50; i <= 140; i++ {
		// string(i) is converting i to a string (i continues to be an integer)

		// []byte(string(i)) is converting the string(i) to a slice of bytes

		// Basically, the []byte(variable) will convert every single character that
		// is part of the variable to the ASCII decimal representation of them
		// and append each converted character to a slice
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}

	hello := "Hello"
	fmt.Println([]byte(hello))
	// [72 101 108 108 111]
	// H   e   l   l   o
	// The slice of bytes holds the decimals representation for each of these
	// characters from the ASCII table
	// If them are converted to string, the slice will represent Hello
	fmt.Println(string([]byte(hello)))

	// Converting a character of a string to integer is the same to convert it to a rune
	// The decimal representation for the character
	// The code below will print the rune that represents H
	fmt.Println(int(hello[0]))
}
