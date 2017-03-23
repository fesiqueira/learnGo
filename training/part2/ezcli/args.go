package ezcli

import (
	"fmt"
	"os"
)

// UserArgs receives the number of maximum arguments the user can input
// and returns a slice of strings containing all the input arguments
// If the CLI is limited to 5 arguments and the user input 4 or less
// the missing arguments will be initilized with empty
func UserArgs(nArgs int) (arguments []string) {
	var args = make([]string, nArgs)
	osArgs := os.Args[1:]  // Initializing a variable that holds only the arguments (except the program location)
	numArgs := len(osArgs) // Initializing a variable that holds the total of arguments passed
	switch {               // Limiting the number of maximum arguments
	case numArgs == 0:
		fmt.Println("Too few arguments, use -h")
	case numArgs > nArgs:
		fmt.Println("Too much arguments, use -h")
	default:
		copy(args, osArgs)
	}
	return args
}
