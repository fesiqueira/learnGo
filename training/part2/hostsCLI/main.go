package main

import (
	"fmt"
	"os"
)

func optionAndHost() (string, string) {
	var args [2]string

	// Initializing a variable that holds only the arguments (except the program location)
	osArgs := os.Args[1:]

	// Initializing a variable that holds the total of arguments passed
	numArgs := len(osArgs)

	// Limiting to 2 arguments
	switch {
	case numArgs == 0:
		fmt.Println("Too few arguments, use -h")
	case len(os.Args[1:]) > 2:
		fmt.Println("Too much arguments, use -h")
	default:
		for i := 0; i < numArgs; i++ {
			args[i] = osArgs[i]
		}
	}
	return args[0], args[1]
}

func main() {
	fmt.Println(optionAndHost())
}
