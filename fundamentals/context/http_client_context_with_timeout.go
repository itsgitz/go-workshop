package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// define the api url or endpoint
	const APIURL string = "http://localhost:8000/players"

	fmt.Println("THE HTTP CLIENT CONTEXT WITHTIMEOUT WORKSHOP", time.Now())

	// define the parent context
	ctx := context.Background()
	fetch(ctx, APIURL)
}

func fetch(ctx context.Context, url string) {
	client := &http.Client{}
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req.WithContext(ctx))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer res.Body.Close()

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(buff))
}
