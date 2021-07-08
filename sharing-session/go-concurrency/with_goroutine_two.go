package main

import (
	"fmt"
	"time"
)

func compute(val int) {
	for i := 0; i < val; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
func main() {
	go compute(10)
	go compute(5)

	fmt.Scanln()
}
