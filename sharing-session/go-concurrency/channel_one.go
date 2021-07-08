package main

import (
	"fmt"
	"time"
)

func main() {
	timeChan := make(chan time.Time)

	go func() {
		time.Sleep(time.Second)
		timeChan <- time.Now()
	}()

	completed := <-timeChan

	fmt.Println(completed)
}
