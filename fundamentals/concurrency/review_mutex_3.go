package main

import (
	"fmt"
	"sync"
)

var (
	m sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	// var memory = []string{"a", "b", "c", "d", "e"}

	fmt.Println("New Go Mutex Review: Ch 3")

	// wg.Add(1)
	// go func(wg *sync.WaitGroup) {
	// 	m.Lock()
	// 	fmt.Println("I'm in concurrency")
	// 	m.Unlock()

	// 	wg.Done()
	// }(&wg)

	// wg.Wait()

	// we will face the race condition here if we don't use the mutex mechanism
	var x = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(w *sync.WaitGroup) {
			defer w.Done()
			m.Lock()
			x = x + 1
			m.Unlock()
		}(&wg)
	}

	wg.Wait()
	fmt.Println(x)
}
