package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func userArgs() (option string, hostname string) {
	var args [2]string
	osArgs := os.Args[1:]  // Initializing a variable that holds only the arguments (except the program location)
	numArgs := len(osArgs) // Initializing a variable that holds the total of arguments passed
	switch {               // Limiting to 2 arguments
	case numArgs == 0:
		fmt.Println("Too few arguments, use -h")
	case numArgs > 2:
		fmt.Println("Too much arguments, use -h")
	default:
		copy(args[:], osArgs)
	}
	return args[0], args[1]
}

func printHelp() {
	fmt.Println("Help log")
}

func main() {
	option, _ := userArgs()
	switch option {
	case "-l":
		fileBytes, err := ioutil.ReadFile("./servers.txt")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(fileBytes))
	case "-h":
		printHelp()
	}
}
