package main

import "fmt"

func main() {
	fibonacci()
}

func fibonacci() {
	fmt.Println("Another Fibonacci Example!")

	var result int
	first := 0
	two := 1
	numberOfSequence := 10

	for i := 0; i < numberOfSequence; i++ {
		fmt.Printf("#%d = %d \n", i, first)
		result = first + two
		first = two
		two = result
	}
}
