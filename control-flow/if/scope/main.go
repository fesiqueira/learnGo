package main

import "fmt"

func main() {
	// A variable is initialized with "Felipe" and right after the semicolon, the if
	// will evaluate if the initialized is identical "Felipe"
	// If true it will run

	// The name scope is only inside the if's curlies, so it's not accessible from anywhere
	// except within the if statement
	if name := "Felipe"; name == "Felipe" {
		fmt.Println(name)
	}
}
