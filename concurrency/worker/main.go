package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	Start(ch)

	go Worker(ch)

	time.Sleep(5 * time.Second)

	Stop(ch)

	var input string
	fmt.Scanln(&input)
}

func Start(ch chan bool) {
	fmt.Println("Work!")
	go func() {
		ch <- true
	}()
}

func Stop(ch chan bool) {
	fmt.Println("Worker, stop!")
	go func() {
		ch <- false
	}()
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
