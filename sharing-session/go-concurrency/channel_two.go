package main

import "fmt"

func calSquare(sOne, sTwo float32, v chan float32) {
	result := sOne * sTwo

	v <- result
}

func calRectangle(p, l float32, v chan float32) {
	result := p * l

	v <- result
}

func main() {
	fmt.Println("Program started")

	squareChan := make(chan float32)
	rectangleChan := make(chan float32)

	go calSquare(5, 5, squareChan)
	go calRectangle(8, 6, rectangleChan)

	getSquare := <-squareChan
	getRectangle := <-rectangleChan

	fmt.Println("square:", getSquare)
	fmt.Println("rectangle:", getRectangle)

	fmt.Println("Finish")
}
