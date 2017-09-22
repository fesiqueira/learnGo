package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Worker struct {
	Name string
	Chan chan bool
	Task func()
}

type WorkerJSON struct {
	Name string `json:"workername"`
	Task string `json:"task"`
}

func NewWorker(name string, task func(workerName string, ch chan bool) func()) *Worker {
	w := &Worker{
		Name: name,
		Chan: make(chan bool),
	}

	w.Task = task(w.Name, w.Chan)

	return w
}

func (w *Worker) Work() {
	fmt.Println(w.Name, "is working")
	go w.Task()
}

func (w *Worker) Stop() {
	fmt.Println(w.Name, "stop")
	go func() {
		w.Chan <- false
	}()
}

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

var workers = map[string]*Worker{
	"worker1": NewWorker("Worker 1", Increment),
	"worker2": NewWorker("Worker 2", Increment),
	"worker3": NewWorker("Worker 3", Increment),
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/workers", ListWorkers)
	router.HandleFunc("/start/{workername}", Start(workers))
	router.HandleFunc("/stop/{workername}", Stop(workers))

	http.ListenAndServe(":4000", router)
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

func ListWorkers(w http.ResponseWriter, r *http.Request) {
	workersJSON := make([]WorkerJSON, 0)

	for _, wrkr := range workers {
		workersJSON = append(workersJSON, WorkerJSON{wrkr.Name, GetTaskName(wrkr.Task)})
	}

	encod := json.NewEncoder(w)
	encod.SetIndent("", " ")

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := encod.Encode(workersJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetTaskName(i interface{}) string {
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), ".")[1]
}
