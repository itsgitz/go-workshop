package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutine using Select")

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case p1 := <-c1:
			fmt.Println("Received:", p1)
		case p2 := <-c2:
			fmt.Println("Received:", p2)
		}
	}

	// if we make the loop operation greater than the number of channel or we loop it forever, the 'deadlock' error will appear
	// for {
	// 	select {
	// 	case p1 := <-c1:
	// 		fmt.Println("Received:", p1)
	// 	case p2 := <-c2:
	// 		fmt.Println("Received:", p2)
	// 	}
	// }
}
