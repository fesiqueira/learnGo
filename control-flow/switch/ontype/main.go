package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // This is assert; asserting "x is of this type"
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("I don't know")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Felipe")
	SwitchOnType(SwitchOnType)
	var t = Contact{"Hello", "Felipe"}
	fmt.Println(t)
	SwitchOnType(t)
}
