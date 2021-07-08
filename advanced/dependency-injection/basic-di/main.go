package main

import "fmt"

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
	User    User
}

type User struct {
	Name string
}

func NewMessage() Message {
	return Message("Hi There!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func (u User) ShowUser() string {
	return u.Name
}

func NewEvent(g Greeter, u User) Event {
	return Event{Greeter: g, User: u}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
	fmt.Println("My name is", e.User.ShowUser())
}

func NewUser() User {
	return User{Name: "Anggit"}
}

func main() {
	e := InitializeEvent()
	e.Start()
}
