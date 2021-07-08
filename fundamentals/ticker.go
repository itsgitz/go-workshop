package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Go Ticker")

	tick := time.NewTicker(3 * time.Second)

	for _ = range tick.C {
		fmt.Println("For tock ...")
	}
}
