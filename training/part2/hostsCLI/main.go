package main

import (
	"fmt"
	"os"
)

func main() {
	var osArgs [2]string

	if len(os.Args[1:]) == 0 {
		fmt.Println("Too few arguments, use -h")
		return
	} else if len(os.Args[1:]) > 2 {
		fmt.Println("Too much arguments, use -h")
		return
	}

	for i := 0; i < len(os.Args[1:]); i++ {
		osArgs[i] = os.Args[i+1]
	}

	fmt.Println(osArgs)
}
