package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("It's even", i)
		} else {
			fmt.Println("It's odd", i)
		}
	}
}
