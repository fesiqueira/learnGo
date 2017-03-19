package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Bom dia!",
	}

	// Range has 2 returns, the first is the Key (or the index, for arrays/slices)
	// The second is the value
	// The for with range will iterate over every item (for each)
	for key, val := range myMap {
		fmt.Println(key, "-", val)
	}
}
