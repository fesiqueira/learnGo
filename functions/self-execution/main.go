package main

import "fmt"

func main() {
	func() int {
		x++
		return x
	}()
	fmt.Println(x)
}

var x int
