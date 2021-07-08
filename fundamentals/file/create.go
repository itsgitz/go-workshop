package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	fileName = "write.txt"
)

var (
	firstContent  = "First Content \n\n"
	secondContent = "Second Content \n\n"
)

func main() {
	err := ioutil.WriteFile(fileName, []byte(firstContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	write, err := file.WriteString(secondContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("write:", write)
}
