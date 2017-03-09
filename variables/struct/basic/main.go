package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) printAge() {
	fmt.Println(u.age)
}

func main() {
	user1 := User{"Felipe", 20}

	user1.printAge()
}
