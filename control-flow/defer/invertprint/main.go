package main

import "fmt"

func main() {
	for i := 1; i < 6; i++ {
		defer fmt.Printf("%v", i)
	}
}
