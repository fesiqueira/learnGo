package main

import (
	"fmt"
	"log"
)

func main() {
	arduinos := FindArduinos()

	for name, arduino := range arduinos {
		err := arduino.AddDevice("led", "13")
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf(
			"Name: %s\nPort: %s\nDevices: %d\nStatus: %t\nLEDs: %d\n",
			name,
			arduino.Port,
			arduino.Devices.CountTotal(),
			arduino.Status,
			arduino.Devices.CountPerDevice()["LED"],
		)
	}
}
