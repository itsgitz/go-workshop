package main

import "fmt"

// Service interface
type Service interface {
	SayHi()
}

// FirstService struct
type FirstService struct{}

// SayHi methods for FirstService struct
func (s FirstService) SayHi() {
	fmt.Println("Hi from FirstService!")
}

// SecondService struct
type SecondService struct{}

// SayHi method for SecondService struct
func (s SecondService) SayHi() {
	fmt.Println("Hi from SecondService!")
}

func main() {
	fmt.Println("Go Map Workshop")

	// we can define a map of string uuids to the interface type`Service`
	i := make(map[string]Service)

	// we can then populate our map with simple ids to particular services
	i["SERVICE_ID_1"] = FirstService{}
	i["SERVICE_ID_2"] = SecondService{}

	// incoming http request wants service 2
	// we can use the incoming uuid to lookup the required
	// service and call it's SayHi() method
	// i["SERVICE_ID_2"].SayHi()

	runService(i)
}

func runService(service map[string]Service) {
	for k, s := range service {
		fmt.Println("Service:", k)
		s.SayHi()
	}
}
