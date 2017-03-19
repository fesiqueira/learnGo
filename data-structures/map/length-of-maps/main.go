package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["Felipe"] = "Siqueira"
	myMap["Pinheiro"] = "Siqueira"
	myMap["Siqueira"] = "Siqueira"

	fmt.Println(myMap)
	fmt.Println(len(myMap))

}
