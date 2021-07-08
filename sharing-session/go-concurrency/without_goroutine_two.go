package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}

func sayMyName(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name)
	}
}

func main() {
	name := "udin"
	fmt.Println("A simple program without goroutine")
	fmt.Println("Program started ...")

	go sayHello()
	sayHello()
	sayMyName(name)

	fmt.Println("Finished!")

	time.Sleep(3 * time.Second)
}
