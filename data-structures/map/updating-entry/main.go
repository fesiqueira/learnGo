package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["Felipe"] = "Siqueira"

	// Updates the value for the "Felipe" key
	myMap["Felipe"] = "Pinheiro"

	fmt.Println(myMap)
}
