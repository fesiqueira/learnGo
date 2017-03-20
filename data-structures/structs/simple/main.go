// In other languages the Classes are the blueprints that are used to instantiate new objects

// In Go, Types are the blueprints, and the objects are just variables of the created type

package main

import "fmt"

// We created a type called person, and it's underlying type is a struct, that carries three fields (variables) within it
// 2 strings and a int
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Felipe", "Siqueira", 20}
	p2 := person{"Solange", "de Camargo", 44}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
