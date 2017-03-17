package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func showServers() string {
	serverFile := "./servers.txt"

	fileContent, err := ioutil.ReadFile(serverFile)
	checkError(err)

	return string(fileContent)
}

// note, that variables are pointers
var intFlag = flag.Int("port", 0, "Flag usage example: port")
var listFlag = flag.String("l", showServers(), "Used to show all available servers")

func main() {
	flag.Parse()
	fmt.Println(*listFlag)
	fmt.Println("Server running on port:", *intFlag)
	fmt.Println(flag.Args())
}
