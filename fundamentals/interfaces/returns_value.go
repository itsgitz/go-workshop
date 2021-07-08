package main

import "fmt"

func main() {
	var programmers []Employee

	gitz := Engineer{Name: "Gitz", Age: 23}
	ryouko := Engineer{Name: "Ryouko", Age: 24}
	programmers = append(programmers, gitz)
	programmers = append(programmers, ryouko)

	fmt.Println("programmers:", programmers)
}

type Employee interface {
	SetLanguage() string
	SetAge() int
}

type Engineer struct {
	Name string
	Age  int
}

func (e Engineer) SetLanguage() string {
	return e.Name + " programs in Go"
}

func (e Engineer) SetAge() int {
	return e.Age
}
