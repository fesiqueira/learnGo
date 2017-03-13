package main

import "fmt"

func main() {
	divideAndEven := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	fmt.Println(divideAndEven(1))
	fmt.Println(divideAndEven(10))
	fmt.Println(divideAndEven(531))
}
