package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			// Break will break completelly the loop, putting the program to run after the loop
			break
		}
		i++
	}
}
