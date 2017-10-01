package main

import (
	"io/ioutil"
	"strings"
)

// Arduino represents a physical Arduino hardware
type Arduino struct {
	Port string // port which the Arduino is connected
}

// findArduinos will iterate over all items found in /dev, and append
// the ones that names is prefixed by ttyUSB (commonly Arduinos)
func findArduinos() ([]Arduino, error) {
	var (
		arduinos      = make([]Arduino, 0)     // initialing the slice of Arduinos
		contents, err = ioutil.ReadDir("/dev") // reading all items within /dev directory
	)

	// if something went wrong while reading the items, return the empty slice of arduinos and the error
	if err != nil {
		return arduinos, err
	}

	// iterate over every item
	for _, content := range contents {
		// get the name of every item, and check if it's prefixed with ttyUSB
		if contentName := content.Name(); strings.HasPrefix(contentName, "ttyUSB") {
			arduinos = append(arduinos, Arduino{contentName}) // append the item to the slice of Arduinos
		}
	}

	return arduinos, nil
}
