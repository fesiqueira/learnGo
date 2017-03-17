package main

import (
	"flag"
	"fmt"
)

// note, that variables are pointers
var intFlag = flag.Int("port", 0, "Flag usage example: port")
var listFlag = flag.String("l", showServers(), "Used to show all available servers")

func main() {
	flag.Parse()
	fmt.Println("Server running on port:", *intFlag)
}
