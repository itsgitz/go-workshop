package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiEndpoint string = "http://localhost:8000/players"

func main() {
	fetch(apiEndpoint)
}

func fetch(url string) {
	c := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(buff))
}
