package main

import (
	"errors"
	"reflect"
	"strconv"

	"gobot.io/x/gobot/drivers/gpio"

	"gobot.io/x/gobot/platforms/firmata"
)

// Devices represents all devices that can be attached to an Arduino
type Devices struct {
	LED   map[string]*gpio.LedDriver
	Motor map[string]*gpio.MotorDriver
}

// InitDevices initializes the maps in Devices
func InitDevices() Devices {
	return Devices{
		LED:   make(map[string]*gpio.LedDriver),
		Motor: make(map[string]*gpio.MotorDriver),
	}
}

// Len returns the number of devices attached to an Arduino
func (d *Devices) Len() int {
	var devices int
	value := reflect.ValueOf(*d)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()

		switch reflect.TypeOf(field).Kind() {
		case reflect.Map:
			devices += reflect.ValueOf(field).Len()
		}
	}

	return devices
}

// Arduino represents a physical Arduino
type Arduino struct {
	Name       string
	Port       string
	Connection *firmata.Adaptor
	Devices    Devices
}

// Arduinos represents a slice of pointer of Arduino
type Arduinos []*Arduino

// NewArduino initializes a new Arduino object, based on the serial connection
func NewArduino(number int, port string) *Arduino {
	return &Arduino{
		Name:       "Arduino-" + strconv.Itoa(number),
		Port:       port,
		Connection: firmata.NewAdaptor(port),
		Devices:    InitDevices(),
	}
}

// AddDevice appends a new gobot.Device into Arduino.Devices
func (a *Arduino) AddDevice(deviceType, pin string) error {
	switch deviceType {
	case "led", "Led", "LED":
		ledNumber := strconv.Itoa(len(a.Devices.LED) + 1)
		a.Devices.LED["LED-"+ledNumber] = gpio.NewLedDriver(a.Connection, pin)
	case "motor", "Motor", "MOTOR":
		motorNumber := strconv.Itoa(len(a.Devices.Motor) + 1)
		a.Devices.Motor["MOTOR"+motorNumber] = gpio.NewMotorDriver(a.Connection, pin)
	default:
		return errors.New("Unsupported device")
	}

	return nil
}
