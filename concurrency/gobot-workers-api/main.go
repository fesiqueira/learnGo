package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// arduinos, err := serial.GetPortsList()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// firmataAdaptor := firmata.NewAdaptor("ttyUSB0")
	// led := gpio.NewLedDriver(firmataAdaptor, "13")

	// devices := gobot.Devices{led}
	var workers = map[string]*Worker{
		// "worker1": NewWorker("Worker 1", Increment),
		// "worker2": NewWorker("Worker 2", Increment),
		// "worker3": NewWorker("Worker 3", Increment),
		"bot": NewWorker("bot", ledBlink),
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/workers", ListAllWorkers(workers))
	router.HandleFunc("/start/{workername}", Start(workers))
	router.HandleFunc("/stop/{workername}", Stop(workers))
	router.HandleFunc("/add/{workername}", AddLedToWorker(workers))

	http.ListenAndServe(":4000", router)
}
