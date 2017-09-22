package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	workers := 1
	ch := make(chan bool, workers)

	router := mux.NewRouter().StrictSlash(true)

	go func() {
		ch <- false
	}()

	go Worker(ch)

	router.HandleFunc("/start", Start(ch))
	router.HandleFunc("/stop", Stop(ch))

	http.ListenAndServe(":3000", router)
}

func Worker(ch chan bool) {
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
			time.Sleep(1 * time.Second)
		}
	}
}

func Start(ch chan bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Work!")

		go Worker(ch)
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
