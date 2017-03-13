package main

import "fmt"

func main() {
	divideAndEven := func(x int) (half int, even bool) {
		if x%2 == 0 {
			even = true
		}
		half = x / 2
		return half, even
	}
	fmt.Println(divideAndEven(1))
	fmt.Println(divideAndEven(10))
	fmt.Println(divideAndEven(531))
}
