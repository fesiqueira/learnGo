package main

import "fmt"

func main() {
	greet("Felipe")
	greet("Siqueira")
}

func greet(name string) {
	fmt.Println("Hello", name)
}
