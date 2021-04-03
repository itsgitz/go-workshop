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
	const url string = "http://localhost:8000/players"

	fmt.Println("HTTP CLIENT CONTEXT WITH CANCEL WORKSHOP", time.Now())

	fetch(context.Background(), url)
}

func fetch(ctx context.Context, url string) {
	client := &http.Client{}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

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
