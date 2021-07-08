package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

// JSONContent for bla bla bla
type JSONContent struct {
	MessageID string    `json:"message_id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Date      time.Time `json:"date"`
}

func main() {
	run()
}

func run() {
	fmt.Println("Encode JSON")
	err := writeJSONFile()
	if err != nil {
		log.Fatal(err)
	}
}

func jsonEncode() ([]byte, error) {
	content := &JSONContent{
		MessageID: "123",
		From:      "anggit@isi.co.id",
		To:        "ryoko@heaven.co.id",
		Date:      time.Now(),
	}

	j, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func writeJSONFile() error {
	const filename = "email_pipe.json"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	data, err := jsonEncode()
	if err != nil {
		return err
	}

	write, err := f.Write(data)
	if err != nil {
		return err
	}

	if write == 0 {
		return errors.New("Failed to write jsonfile")
	}

	return nil
}
