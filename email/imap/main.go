package main

import (
	"sync"

	"github.com/itsgitz/go-workshop/email/imap/imap"
)

func main() {
	var i imap.Imap
	var wg sync.WaitGroup

	wg.Add(1)

	i = &imap.Client{}

	i.New()
	i.GetMessageBody(&wg)

	wg.Wait()
}
