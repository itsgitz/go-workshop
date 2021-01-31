package imap

import (
	"log"
	"sync"

	gc "github.com/emersion/go-imap/client"
)

// Imap interface
type Imap interface {
	New()
	GetConnection() *gc.Client
	GetMessages()
	GetMessageBody(wg *sync.WaitGroup)
	CopyEmailToFile()
}

// New method for initialize new imap client
func (cl *Client) New() {
	log.Println("Connecting to IMAP Server ...")

	cl.setCredentials()
	cl.setImapServer()
	cl.login()
}
