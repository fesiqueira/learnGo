package main

import "fmt"

func divideAndEven(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	fmt.Println(divideAndEven(1))
	fmt.Println(divideAndEven(10))
	fmt.Println(divideAndEven(531))
}
