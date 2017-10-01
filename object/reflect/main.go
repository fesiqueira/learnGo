package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Username string
	Age      int
}

func (u *User) MethodOne() {
	fmt.Println(u.Username)
}

func (u *User) MethodTwo() {
	fmt.Println(u.Age)
}

func main() {
	user := reflect.TypeOf(&User{})

	for i := 0; i < user.NumMethod(); i++ {
		method := user.Method(i)
		fmt.Println(method.Type)
	}
}
