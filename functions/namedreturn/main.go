package main

import (
	"fmt"
)

func main() {

	fmt.Println(greet("Felipe", "Siqueira"))
}

func greet(firstname, lastname string) (s string) {
	s = fmt.Sprint(firstname, lastname)
	return
}
