package main

import (
	"log"
	"os"
)

const (
	filename = "write.txt"
)

func main() {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	write, err := f.Write([]byte("Append line \n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Write:", write)
}
