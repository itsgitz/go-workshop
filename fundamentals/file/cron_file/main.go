package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	filename = "some.txt"
)

func main() {
	content := []byte("Ohayou! \n")

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	write, err := f.Write(content)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Write:", write)

	// show the content
	o, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(o))
}
