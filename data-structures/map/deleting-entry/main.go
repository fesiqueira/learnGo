package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Bom dia!",
	}

	fmt.Println(myMap)
	delete(myMap, 2) // Deletes based on the key, in this case delete the key 2
	fmt.Println(myMap)
}
