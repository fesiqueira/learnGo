package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	serial "go.bug.st/serial.v1"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"

	"gobot.io/x/gobot/platforms/firmata"
)

// Arduino represents a physical Arduino
type Arduino struct {
	Name string
	Port string
	Status
	Connection *firmata.Adaptor
	Devices    map[string]gobot.Driver
}

// Status concentrates the status info about an Arduino
type Status struct {
	State     bool
	Connected bool
}

// Arduinos represents a slice of pointer of Arduino
type Arduinos map[string]*Arduino

// NewArduino initializes a new Arduino object, based on the serial connection
func NewArduino(number int, port string) *Arduino {
	return &Arduino{
		Name:       "Arduino-" + strconv.Itoa(number),
		Port:       port,
		Connection: firmata.NewAdaptor(port),
		Devices:    make(map[string]gobot.Driver),
	}
}

// FindArduinos scans all serial ports and initializes all found Arduinos
func FindArduinos() Arduinos {
	var (
		arduinos     = make(Arduinos)
		serials, err = serial.GetPortsList()
	)

	if err != nil {
		log.Fatal(err)
	}

	for number, serial := range serials {
		// if strings.HasPrefix(serial, "ttyUSB") {
		arduino := NewArduino(number, serial)
		arduinos[arduino.Name] = arduino
		// }
	}

	return arduinos
}

// AddDevice appends a new gobot.Device into Arduino.Devices
func (a *Arduino) AddDevice(deviceType, pin string) error {
	switch deviceType {
	case "led", "Led", "LED":
		ledNumber := strconv.Itoa(len(a.Devices) + 1)
		a.Devices["LED-"+ledNumber] = gpio.NewLedDriver(a.Connection, pin)
	case "motor", "Motor", "MOTOR":
		// motorNumber := strconv.Itoa(len(a.Devices.Motor) + 1)
		// a.Devices.Motor["MOTOR"+motorNumber] = gpio.NewMotorDriver(a.Connection, pin)
	default:
		return errors.New("Unsupported device")
	}

	return nil
}

// Start "boots" an Arduino, starting the connection and it's devices
func (a *Arduino) Start() {
	a.Connection.Connect()
	a.Connected = true

	for _, device := range a.Devices {
		device.Start()
	}
}

func (a *Arduino) ListDevices() {
	for name, device := range a.Devices {
		fmt.Printf("Name: %s\nDevice: %p\n\n", name, device)
	}
}
