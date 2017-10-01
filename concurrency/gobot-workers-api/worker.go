package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"gobot.io/x/gobot/platforms/firmata"

	"gobot.io/x/gobot"
)

type taskFunc func(w *Worker) func()

type Worker struct {
	Name       string
	connection *firmata.Adaptor
	devices    gobot.Devices
	order      chan bool
	task       func()
	status     bool
}

func NewWorker(name string, task taskFunc) *Worker {
	w := &Worker{
		Name:       name,
		connection: firmata.NewAdaptor("/dev/ttyUSB0"),
		order:      make(chan bool),
		status:     false,
	}

	w.task = task(w)

	return w
}

func (w *Worker) Work() {
	if w.status {
		fmt.Println(w.Name, "is already working")
		return
	}
	w.status = true
	go w.task()
}

func (w *Worker) Stop() {
	if !w.status {
		fmt.Println(w.Name, "is not running")
		return
	}
	w.status = false
	go func() {
		w.order <- false
	}()
}

func (w *Worker) Connection() *firmata.Adaptor {
	return w.connection
}

func (w *Worker) Devices() []string {
	var devs = make([]string, 0)

	for _, device := range w.devices {
		devs = append(devs, device.Name())
	}

	return devs
}

func (w *Worker) TaskName() string {
	// return strings.Split(strings.Split(runtime.FuncForPC(reflect.ValueOf(w.task).Pointer()).Name(), "/")[4], ".")[1]
	return strings.Split(runtime.FuncForPC(reflect.ValueOf(w.task).Pointer()).Name(), ".")[1]
}

func (w *Worker) Status() string {
	if w.status {
		return "Running"
	}

	return "Idle"
}
