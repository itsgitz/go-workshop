package main

import "fmt"

func main() {
	fibonacci()
}

// goal: 0 1 1 2 3 5 8 13 21 34 55 89 144
func fibonacci() {
	var result [20]int
	numberOfSequence := 20
	startNumber := 0

	// formula:
	// fn = fn-1 + fn-2
	for i := startNumber; i < numberOfSequence; i++ {
		if i <= 1 {
			result[i] = i
		} else {
			result[i] = result[(i-1)] + result[(i-2)]
		}
		fmt.Printf("Sequence #f%d = %d\n", i, result[i])
	}
}
