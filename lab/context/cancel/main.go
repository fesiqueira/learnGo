package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx := context.Background()
	withcancel, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		signal.Stop(c)
		cancel()
	}()

	go CancelableFunction(withcancel, &wg)

	wg.Wait()
}

func CancelableFunction(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			defer wg.Done()
			fmt.Println("Canceled")
			return
		default:
			select {
			case <-ctx.Done():
			}
		}
	}
}
