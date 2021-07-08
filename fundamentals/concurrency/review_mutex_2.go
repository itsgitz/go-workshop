package main

import (
	"fmt"
	"sync"
)

// Anime interface
type Anime interface {
	New(name string, hp int)
	SetPower(power float32)
	GetName() string
	GetPower() float32
	ShowCharacter()
	UpgradePower(percent int, wg *sync.WaitGroup)
	Attacked(damage int, wg *sync.WaitGroup)
}

// OnePiece data collection
type OnePiece struct {
	Mutex sync.Mutex
	Name  string
	Power float32
	Anime string
	HP    int
}

// Naruto data collection
type Naruto struct {
	Name  string
	Power int
	Anime string
}

// New method for set new One Piece character
func (o *OnePiece) New(name string, hp int) {
	o.HP = hp
	o.Name = name
	o.Anime = "One Piece"
}

// SetPower for set the character's power
func (o *OnePiece) SetPower(power float32) {
	o.Power = power
}

// GetName for get the character name
func (o *OnePiece) GetName() string {
	return o.Name
}

// GetPower for get the power of character
func (o *OnePiece) GetPower() float32 {
	return o.Power
}

// ShowCharacter for display the character details
func (o *OnePiece) ShowCharacter() {
	fmt.Println("Detail Character of", o.Anime)
	fmt.Println("Name:", o.Name)
	fmt.Println("Power:", o.Power)
	fmt.Println("HP:", o.HP)
}

// UpgradePower for upgrade the power of character
func (o *OnePiece) UpgradePower(percent int, wg *sync.WaitGroup) {
	o.Mutex.Lock()
	upgradeValue := o.Power * float32(percent) / 100
	o.Power = o.Power + upgradeValue
	o.Mutex.Unlock()

	wg.Done()
}

// Attacked from the enemies
func (o *OnePiece) Attacked(damage int, wg *sync.WaitGroup) {
	o.Mutex.Lock()
	o.HP = o.HP - damage
	o.Power = o.Power - 5
	o.Mutex.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var anime Anime

	wg.Add(2)

	anime = &OnePiece{}
	anime.New("Monkey D Luffy", 100)
	anime.SetPower(90)
	anime.ShowCharacter()

	fmt.Println()
	fmt.Println("Upgrading ... ")
	go anime.UpgradePower(10, &wg)

	fmt.Println("Oh no!! Got attacked by enemy!")
	go anime.Attacked(10, &wg)

	fmt.Println()
	wg.Wait()

	anime.ShowCharacter()
}
