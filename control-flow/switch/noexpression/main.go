package main

import "fmt"

func main() {
	x := 0
	switch {
	case x != 0:
		fmt.Println("It won't run")
	case x == 0:
		fmt.Println("Switch without expression will run the FIRST statement that evaluates to true")
	default:
		fmt.Println("None")
	}
}
