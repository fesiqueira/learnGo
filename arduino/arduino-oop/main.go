package main

import (
	"fmt"
	"log"

	serial "go.bug.st/serial.v1"
)

func main() {
	var arduinos Arduinos

	serials, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	for number, serial := range serials {
		// if strings.HasPrefix(serial, "ttyUSB") {
		arduinos = append(arduinos, NewArduino(number, serial))
		// }
	}

	for _, arduino := range arduinos {
		err := arduino.AddDevice("led", "13")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Default Name:", arduino.Name, "Port:", arduino.Port, "Connection:", arduino.Connection.Name(), "Number of Devices:", arduino.Devices.Len())
	}

}
