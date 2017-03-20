package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// fullName is a method of person
// an "action" that person can do
// (p person) is the receiver, almost like the "this" present on other languages
func (p person) fullName() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"Felipe", "Siqueira", 20}
	p2 := person{"Solange", "de Camargo", 44}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
}
