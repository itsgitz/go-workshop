package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/itsgitz/workshop/ory-kratos/functions"
)

const (
	getStatusAlive = "GetStatusAlive"
	getIdentities  = "GetIdentities"
	createIdentity = "CreateIdentity"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Panicln("Error: argument cannot be empty or less than 2. Usage: kratos [email] [username]")
		os.Exit(1)
	}

	fmt.Println("ORY Kratos service command line")
	fmt.Println("Your command arguments: ", args)

	switch args[1] {
	case getStatusAlive:
		// get status alive
		for !functions.GetStatusAlive() {
			fmt.Println("Waiting for ORY Kratos connection ...")
			time.Sleep(time.Second * 1)
		}

		fmt.Println("Finish, ORY Kratos connection is established :)")

		break
	case getIdentities:
		// get all identities function
		functions.GetIdentities()
		break
	case createIdentity:
		// create identity function
		functions.CreateIdentity(args)
		break
	}
}
