package main

import (
	"fmt"
	"os"
)

func userArgs(nArgs int) (arguments []string) {
	var args = make([]string, nArgs)
	osArgs := os.Args[1:]  // Initializing a variable that holds only the arguments (except the program location)
	numArgs := len(osArgs) // Initializing a variable that holds the total of arguments passed
	switch {               // Limiting to 2 arguments
	case numArgs == 0:
		fmt.Println("Too few arguments, use -h")
	case numArgs > nArgs:
		fmt.Println("Too much arguments, use -h")
	default:
		copy(args[:], osArgs)
	}
	return args
}

func main() {
	arguments := userArgs(2)
	fmt.Println(arguments)
}
