package imap

import (
	"log"
	"sync"

	gc "github.com/emersion/go-imap/client"
)

// Client for bind imap connection
type Client struct {
	Mutex sync.Mutex
	*gc.Client
	Credentials
	Mailbox
}

// Login for login to imap server, return imap connection
func (cl *Client) login() {
	if cl.Credentials.ImapTLS {
		c, err := gc.DialTLS(cl.Credentials.ImapServer, nil)
		if err != nil {
			log.Fatal(err)
		}

		defer c.LoggedOut()

		log.Println("Connected to IMAP Server ...")

		if err := c.Login(cl.Credentials.ImapUsername, cl.Credentials.ImapPassword); err != nil {
			log.Fatal(err)
		}

		log.Println("Logged in!")

		cl.Client = c
	}

	c, err := gc.Dial(cl.Credentials.ImapServer)
	if err != nil {
		log.Fatal(err)
	}

	defer c.LoggedOut()

	log.Println("Connected to IMAP Server ...")

	if err := c.Login(cl.Credentials.ImapUsername, cl.Credentials.ImapPassword); err != nil {
		log.Fatal(err)
	}

	log.Println("Logged in!")

	cl.Client = c
}

// GetConnection for get current imap connection
func (cl *Client) GetConnection() *gc.Client {
	return cl.Client
}
