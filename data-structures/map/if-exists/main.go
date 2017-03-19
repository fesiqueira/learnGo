package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Bom dia!",
	}

	fmt.Println(myMap)

	// The val, exists := myMap[index] is called "comma ok" idiom
	// If myMap[2] exists, the val will be set to the value of the key 2, and exists will be TRUE
	// Otherwise, the val will be "zero value", and the exists will be FALSE
	if val, exists := myMap[2]; exists {
		delete(myMap, 2)
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val:", val)
		fmt.Println("exists:", exists)
	}

	fmt.Println(myMap)
}
