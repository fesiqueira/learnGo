package main

import "fmt"

func main() {
	// Under a slice there's an underlying array, but
	// under a map there's a HASH TABLE

	// Map is a key:value data-structure
	// The map below: Key: String Type, Value: String type
	var myMap = make(map[string]string)
	myMap["Felipe"] = "Siqueira"
	myMap["Siqueira"] = "Pinheiro"

	fmt.Println(myMap)
}
