package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Felipe", "Siqueira", 20, 001}

	// Marshal returns the struct "converted" into a JSON, containing all the EXPORTED fields, in the format of a []uint8 (a slice of bytes)
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)

	// It's a slice of bytes ([]uint8)
	fmt.Printf("%T\n", bs)

	// Converting the slice of bytes to a string, the result is the JSON, converted from a struct, containing all the EXPORTED fields
	fmt.Println(string(bs))
}
