package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	fmt.Println("OAuth 2.0 with Client Credentials")

	c := clientcredentials.Config{
		ClientID:     "3ef797fa-f1a3-4d88-a70b-140e6ea6c974",
		ClientSecret: "RVYb91AiTH0hysUiWyxmbMsEQ0",
		TokenURL:     "http://localhost:4444/oauth2/token",
	}

	fmt.Println("Config:")
	fmt.Println(c)

	ctx := context.Background()
	token, err := c.Token(ctx)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("token:")
	fmt.Println(token)

	fmt.Println("Get access token:", token.AccessToken)
	// oauth2Client := c.Client(context.Background())

	// fmt.Println("oath2Client:", oauth2Client)

	// action!
}
