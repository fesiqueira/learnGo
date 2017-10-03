package main

import (
	"log"
)

func main() {
	arduinos := FindArduinos()

	for _, arduino := range arduinos {
		err := arduino.AddDevice("led", "13")
		if err != nil {
			log.Fatal(err.Error())
		}

		// fmt.Printf(
		// 	"Name: %s\nPort: %s\nDevices: %s\nStatus: %t\n\n",
		// 	name,
		// 	arduino.Port,
		// 	// arduino.Devices.CountTotal(),
		// 	arduino.Devices,
		// 	arduino.Status,
		// 	// arduino.Devices.CountPerDevice()["LED"],
		// )
		// fmt.Println(arduino.Connected)

		// arduino.Start()

		// fmt.Println(arduino.Connected)

		// arduino.Devices["LED-1"].(*gpio.LedDriver).On()
		arduino.ListDevices()
	}
	// arduinos["Arduino-0"].Devices["LED-1"].(*gpio.LedDriver).On()
	// arduinos["Arduino-1"].Devices["LED-1"].(*gpio.LedDriver).On()
}
