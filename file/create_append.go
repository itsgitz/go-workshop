package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	// define the filename
	filename = "create_append.txt"
)

func main() {
	// define the content
	content := []byte("Anggit and Ryoko forever :) \n")

	// open the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// write the file with content
	write, err := f.Write(content)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Write:", write)

	// print the file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
