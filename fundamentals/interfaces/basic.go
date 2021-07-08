// source: https://tutorialedge.net/golang/go-interfaces-tutorial/
package main

import "fmt"

func main() {
	var bassist BaseGuitarist
	bassist.Name = "Ryouko"
	bassist.PlayGuitar()

	var accoustic AccousticGuitarist
	accoustic.Name = "Anggit"
	accoustic.PlayGuitar()

	var guitarists []Guitarist

	guitarists = append(guitarists, bassist)
	guitarists = append(guitarists, accoustic)

	ShowGuitarist(guitarists)
}

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AccousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the bass guitar\n", b.Name)
}

func (a AccousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the accoustic guitar\n", a.Name)
}

func ShowGuitarist(g []Guitarist) {
	fmt.Println("Our guitarists:", g)
}
