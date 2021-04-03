// This source code is the example of WithTimeout usage insipired by post from linkedin
// https://www.linkedin.com/pulse/its-time-understand-golang-contexts-lucas-schenkel-schieferdecker/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// create child context using WithCancel method, and Background() used as parent context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// cancel the context after five seconds
	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	go doSomething(ctx, 3)
	go loop(ctx, 8)

	fmt.Scanln()
}

func doSomething(ctx context.Context, timeSleep time.Duration) {
	fmt.Println("Starting doSomething ...")
	select {
	case <-time.After(timeSleep * time.Second):
		fmt.Println("Finished doSomething")
	case <-ctx.Done():
		fmt.Println("Cancelling doSomething")
	}
}

func loop(ctx context.Context, n int) {
	for i := 0; i < n; i++ {
		if ctx.Err() != nil {
			fmt.Println("Leaving loop ...")
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Loop #", i)
	}
}
