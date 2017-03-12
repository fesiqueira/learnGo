package main

import (
	"fmt"
)

func main() {
	ret := greet("Felipe", "Siqueira")
	fmt.Println(ret)
}

func greet(name, lastname string) string {
	return fmt.Sprint(name, lastname)
}
