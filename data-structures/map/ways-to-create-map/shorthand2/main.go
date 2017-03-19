package main

import "fmt"

func main() {
	// Map is a key:value data-structure
	// The map below: Key: String Type, Value: String type

	// Initialized a map with composite literals
	myMap := map[string]string{
		"Felipe": "Siqueira",
	}

	myMap["Siqueira"] = "Pinheiro"

	fmt.Println(myMap)
}
