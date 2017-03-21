// [_Command-line arguments_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// are a common way to parameterize execution of programs.
// For example, `go run hello.go` uses `run` and
// `hello.go` arguments to the `go` program.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func optionAndServer() (string, string) {
}

func showAndFind(serverFile, server string) {
	fileBytes, err := ioutil.ReadFile(serverFile)
	checkError(err)

	fmt.Println(string(fileBytes))
}

func main() {
	serverFile := "./servers.txt"

	option, server := optionAndServer()

	switch option {
	case "-l":
		showAndFind(serverFile, server)
	}
	fmt.Println(serverFile)
	fmt.Println(option)
	fmt.Println(server)
}
