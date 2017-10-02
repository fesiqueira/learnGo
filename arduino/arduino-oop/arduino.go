package main

import (
	"errors"
	"strconv"

	"gobot.io/x/gobot/drivers/gpio"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/firmata"
)

type Arduino struct {
	Name       string
	Port       string
	Connection *firmata.Adaptor
	Devices    gobot.Devices
}

type Arduinos []*Arduino

func NewArduino(number int, port string) *Arduino {
	return &Arduino{
		Name:       "Arduino-" + strconv.Itoa(number),
		Port:       port,
		Connection: firmata.NewAdaptor(port),
	}
}

func (a *Arduino) AddDevice(kind, pin string) error {
	switch kind {
	case "Led", "LED", "led":
		device := gpio.NewLedDriver(a.Connection, pin)
		a.Devices = append(a.Devices, device)
	default:
		return errors.New("Unsupported device")
	}

	return nil
}
