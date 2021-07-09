package main

import (
	"fmt"
	"time"
)

// main goroutine, go1
func main() {
	fmt.Println("Goroutine")
	fmt.Println("Start")

	name := "udin"

	// go2
	go func(name string) {
		fmt.Println("Nama saya", name)
	}(name)

	// go3
	/// .....

	fmt.Println("Finish")

	time.Sleep(3 * time.Second)
}
