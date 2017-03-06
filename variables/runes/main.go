package main

import "fmt"

func main() {
	myname := "Felipe"
	myrunes := []byte(myname)
	fmt.Println("String:", myname)

	// Converting myname to a slice of bytes
	// Each of these bytes are the decimal representation of the characters in the ASCII table
	fmt.Println("Slice of bytes:", []byte(myname))

	// Converting the first character (F) to a RUNE (alias to int)
	// A rune is the integer that represents a character in the ASCII table
	fmt.Println("F to a rune:", int(myname[0]))

	fmt.Println("Slice of runes:", myrunes)

	// Returning the slice of bytes to string
	// Each of the runes will be converted to string and printed
	// Together they will print myname
	fmt.Println("Rune to string:", string(myrunes))
}
