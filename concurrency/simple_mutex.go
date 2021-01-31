package main

import (
	"fmt"
	"sync"
)

// Operation struct data collection
type Operation struct {
	mutex sync.Mutex
	Multi int
	Value int
}

// New method for define the value and how many times user want to add.
func (o *Operation) New(val, times int) {
	o.Value = val
	o.Multi = times
}

// Add value x times
func (o *Operation) Add() {
	o.Value++
}

// Run the operation
func (o *Operation) Run() {
	// The operation below will cause race condition in goroutine or concurrency application.
	// In other words, the result is unexpected.
	// for i := 0; i < o.Multi; i++ {
	// 	go func() {
	// 		o.Add()
	// 	}()
	// }

	// The solution to resolve this issue is using mutex or mutual exclusion mechanism
	// Go has provided the library for mutex operation in `sync` library
	var wg sync.WaitGroup

	for i := 0; i < o.Multi; i++ {
		wg.Add(1)

		go func() {
			o.mutex.Lock()
			o.Add()
			o.mutex.Unlock()

			wg.Done()
		}()
	}

	wg.Wait()
}

// Show or print method for get the result of value after operations
func (o *Operation) Show() {
	fmt.Println("Result:", o.Value)
}

func main() {
	o := &Operation{}

	o.New(10, 50)
	o.Run()
	o.Show()
}
