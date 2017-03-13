package main

import "fmt"

type User struct {
	name    string
	email   string
	age     int
	country string
	city    string
}

func NewUser(name string, age int) *User {
	u := new(User)
	u.name = name
	u.age = age
	return u
}

func main() {
	obj1 := NewUser("Felipe", 20)
	fmt.Println("Obj1:", obj1)
}
