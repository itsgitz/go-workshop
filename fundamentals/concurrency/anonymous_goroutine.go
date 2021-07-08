package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Executing my concurrent anonymous function ...")
	}()

	fmt.Scanln()
}
