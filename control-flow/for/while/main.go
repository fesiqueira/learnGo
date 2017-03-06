package main

import "fmt"

func main() {
	i := 0

	// Basically it says: while i is lesser than 10, print i and increment
	// It's the while-ish for (while loop in Go)
	for i < 10 {
		fmt.Println(i)
		i++
	}
}
