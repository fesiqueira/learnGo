package main

import "fmt"

func main() {
	for i := 65; i < 123; i++ {
		fmt.Println(i, "-", string(i), "-", i%12)
	}
}
