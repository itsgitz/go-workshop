// This source code is the example of WithTimeout usage insipired by post from linkedin
// https://www.linkedin.com/pulse/its-time-understand-golang-contexts-lucas-schenkel-schieferdecker/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// create child context with timeout using parent context from Background() method
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go doSomething(ctx, 5, ch)

	select {
	// I'm waiting to check if the context has timed out
	case <-ctx.Done():
		fmt.Println("Context cancelled:", ctx.Err())
	case result := <-ch:
		fmt.Println("doSomething finished", result)
	}
}

func doSomething(ctx context.Context, timeSleep time.Duration, ch chan string) {
	fmt.Println("Sleeping")
	time.Sleep(timeSleep * time.Second)

	fmt.Println("Waking up ...")
	ch <- "Did something!"
}
