package main

import "fmt"

func main() {
	x := 100
	y := 10

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	p := &x
	q := &y

	fmt.Println("*p:", *p)
	fmt.Println("*q:", *q)

	*p = 69
	*q = *p + 1
	fmt.Println("x after pointed:", x)
	fmt.Println("y after pointed:", y)
}
