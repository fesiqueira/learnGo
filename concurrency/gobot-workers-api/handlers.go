package main

import (
	"encoding/json"
	"net/http"

	"gobot.io/x/gobot/drivers/gpio"

	"github.com/gorilla/mux"
)

type WorkerJSON struct {
	Name    string   `json:"name"`
	Devices []string `json:"devices"`
	Task    string   `json:"task"`
	Status  string   `json:"status"`
}

func Start(workers map[string]*Worker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workername := vars["workername"]

		workers[workername].Work()
	})
}

func Stop(workers map[string]*Worker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workername := vars["workername"]

		workers[workername].Stop()
	})
}

func ListAllWorkers(workers map[string]*Worker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		workersJSON := make([]WorkerJSON, 0)

		for _, wrkr := range workers {
			workersJSON = append(
				workersJSON,
				WorkerJSON{
					wrkr.Name,
					wrkr.Devices(),
					wrkr.TaskName(),
					wrkr.Status(),
				},
			)
		}

		encodr := json.NewEncoder(w)
		encodr.SetIndent("", " ")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		err := encodr.Encode(workersJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func AddLedToWorker(workers map[string]*Worker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workername := vars["workername"]

		led := gpio.NewLedDriver(workers[workername].Connection(), "13")
		workers[workername].devices = append(workers[workername].devices, led)
	})
}
