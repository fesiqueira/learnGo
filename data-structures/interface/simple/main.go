package main

import "fmt"

type Person interface {
	Speak()
	Name() string
	Age() int
}

type Student struct {
	name string
	age  int
}

func (s Student) Age() int {
	return s.age
}

func (s Student) Name() string {
	return s.name
}

func (s Student) Speak() {
	fmt.Println("Hello, my name is", s.Name())
}

type School struct {
	name   string
	people []Person
}

func (s *School) numOfPeople() {
	fmt.Println(len(s.people))
}

func (s *School) addPerson(p Person) []Person {
	s.people = append(s.people, p)
	return s.people
}

func main() {
	s1 := Student{"Felipe", 20}

	school := School{"FATEC", []Person{}}
	school.addPerson(s1)

	fmt.Println(school)
}
