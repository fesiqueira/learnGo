package main

import "fmt"

func main() {
	// Letter will be a int32 (a rune), the decimal representation for capital A on the ASCII table
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T\n", letter)

	// Converting the int32 to string, it will print A
	fmt.Println(string(letter))
}
