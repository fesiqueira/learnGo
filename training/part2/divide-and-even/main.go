package main

import "fmt"

func divideAndEven(x int) (half int, even bool) {
	if x%2 == 0 {
		even = true
	}
	half = x / 2
	return half, even
}

func main() {
	fmt.Println(divideAndEven(1))
	fmt.Println(divideAndEven(10))
	fmt.Println(divideAndEven(531))
}
