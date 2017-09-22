package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	firmataAdaptor := firmata.NewAdaptor("/dev/ttyUSB0")
	led := gpio.NewLedDriver(firmataAdaptor, "13")

	ch := make(chan bool)

	work := func(ch chan bool) func() {
		return func() {
			var i int
			for {
				select {
				case x := <-ch:
					if !x {
						return
					}
				default:
					fmt.Println("Working...", i+1)
					i++
					led.Toggle()
					time.Sleep(1 * time.Second)
				}
			}
		}
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work(ch),
	)

	robot.Connections().Start()
	robot.Devices().Start()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/start", Start(robot))
	router.HandleFunc("/stop", Stop(ch))

	http.ListenAndServe(":3000", router)
}

func Start(robot *gobot.Robot) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Work!")
		go robot.Work()
	})
}

func Stop(ch chan bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Worker, stop!")
		go func() {
			ch <- false
		}()
	})
}
