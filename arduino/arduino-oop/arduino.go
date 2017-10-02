package main

import (
	"errors"
	"log"
	"reflect"
	"strconv"

	serial "go.bug.st/serial.v1"

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

// CountTotal returns the number of devices attached to an Arduino
func (d *Devices) CountTotal() int {
	var countTotal int
	value := reflect.ValueOf(*d)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()

		switch reflect.TypeOf(field).Kind() {
		case reflect.Map:
			countTotal += reflect.ValueOf(field).Len()
		}
	}

	return countTotal
}

// CountPerDevice returns a map[string]int containing the quantity of devices, classified by it's type
func (d *Devices) CountPerDevice() map[string]int {
	countPerDevice := make(map[string]int)

	value := reflect.ValueOf(*d)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i).Interface()
		fieldType := reflect.TypeOf(field).Kind()

		switch fieldType {
		case reflect.Map:
			fieldName := value.Type().Field(i).Name
			fieldLength := reflect.ValueOf(field).Len()

			if fieldLength > 0 {
				countPerDevice[fieldName] = fieldLength
			}
		}
	}

	return countPerDevice
}

// Arduino represents a physical Arduino
type Arduino struct {
	Name       string
	Port       string
	Status     bool
	Connection *firmata.Adaptor
	Devices    Devices
}

// Arduinos represents a slice of pointer of Arduino
type Arduinos map[string]*Arduino

// NewArduino initializes a new Arduino object, based on the serial connection
func NewArduino(number int, port string) *Arduino {
	return &Arduino{
		Name:       "Arduino-" + strconv.Itoa(number),
		Port:       port,
		Connection: firmata.NewAdaptor(port),
		Devices:    InitDevices(),
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
