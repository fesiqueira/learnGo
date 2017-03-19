package main

import "fmt"

func main() {
	// Map is a key:value data-structure
	// The map below: Key: String Type, Value: String type

	// Initialized the map with make
	myMap := make(map[string]string)
	myMap["Felipe"] = "Siqueira"
	myMap["Siqueira"] = "Pinheiro"

	fmt.Println(myMap)
}
