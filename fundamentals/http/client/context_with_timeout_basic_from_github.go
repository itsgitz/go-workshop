// source: https://gist.github.com/superbrothers/dae0030c151d1f3c24311df77405169b
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiEndpointURL string = "http://localhost:8000/players"

func main() {
	fetch(apiEndpointURL)
}

func fetch(url string) {
	// define the http client
	client := &http.Client{}

	// set new request with given url and "GET" as method
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	// create child context from parent context (req.Context())
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)

	res, err := client.Do(req)
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
