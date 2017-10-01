package main

import (
	"fmt"
	"log"

	serial "go.bug.st/serial.v1"
)

func main() {
	// arduinos, err := findArduinos()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	arduinos, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(arduinos)
}
