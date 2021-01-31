package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	var wg sync.WaitGroup

	val := 0

	for i := 0; i < 1000; i++ {
		// Place wait group add method befire goroutine and set
		// the delta value (parameter) according to number of goroutines
		wg.Add(1)
		go func() {
			m.Lock()
			val++
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Value:", val)
}
