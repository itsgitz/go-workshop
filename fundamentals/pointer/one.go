package main

import "fmt"

func main() {
	i, j := 42, 2701

	fmt.Println("i:", i, "j:", j)

	p := &i // point to i

	fmt.Println("p point to i:", *p) // read i through the pointer

	*p = 21 // set i through the pointer

	fmt.Println("new value of i after set through the p (pointer):", i)

	p = &j // point to j
	fmt.Println("p point to j:", *p)
	*p = *p / 37 // divide j through the pointer

	fmt.Println(j)
}
