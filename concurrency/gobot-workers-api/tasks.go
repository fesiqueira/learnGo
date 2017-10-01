package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gobot.io/x/gobot/drivers/gpio"
)

func Increment(workerName string, ch chan bool) func() {
	return func() {
		var i int
		for {
			select {
			case x := <-ch:
				if !x {
					return
				}
			default:
				fmt.Println(workerName, "Working...", i+1)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}
}

func ledBlink(w *Worker) func() {
	return func() {
		led := &gpio.LedDriver{}
		for _, d := range w.devices {
			if strings.Contains(d.Name(), "LED") {
				led = d.(*gpio.LedDriver)
				err := w.Connection().Connect()
				if err != nil {
					log.Fatal(err.Error())
				}
				led.Start()
			}
		}

		for {
			select {
			case x := <-w.order:
				if !x {
					return
				}
			default:
				time.Sleep(1 * time.Second)
				led.On()
				time.Sleep(1 * time.Second)
				led.Off()
			}
		}
	}
}

func listDevices(w *Worker) func() {
	return func() {
		for {
			select {
			case x := <-w.order:
				if !x {
					return
				}
			default:
				if w.devices.Len() > 0 {
					for _, d := range w.devices {
						if strings.HasPrefix(d.Name(), "LED") {
							fmt.Println(d.Name())
						}
					}
				}
				time.Sleep(1 * time.Second)
			}
		}
	}
}
