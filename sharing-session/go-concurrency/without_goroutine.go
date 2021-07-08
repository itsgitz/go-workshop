package main

import (
	"fmt"
)

func rocketLauncher() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("Launced!\n\n")
}

func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello, I'm ready!")
	}
	fmt.Printf("\n\n")
}

func main() {
	fmt.Println("Program started ...")

	rocketLauncher()
	sayHello()

	fmt.Println("Program finished")
}
