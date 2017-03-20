package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

// Now the DoubleZero has access to any field of Person
// It's kind like: DoubleZero extends Person
// In Go: Person is promoted to DoubleZero (the inner type is promoted to the outer type)
type DoubleZero struct {
	Person
	// As the Person type has a field called First, the field First within DoubleZero is overriding the other field
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		// The First field below is overriding the field that comes with Person
		// so to have access to the Person First, it should be called like Person.First
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	// Fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(p1.First, p1.Person.First)
}
