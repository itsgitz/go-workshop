package imap

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Credentials struct
type Credentials struct {
	ImapHostname string
	ImapServer   string
	ImapPort     uint32
	ImapUsername string
	ImapPassword string
	ImapTLS      bool
}

// SetCredentials for load imap credentials data from environment variables
func (cl *Client) setCredentials() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cl.Credentials.ImapHostname = os.Getenv("IMAP_HOSTNAME")
	cl.Credentials.ImapUsername = os.Getenv("IMAP_USERNAME")
	cl.Credentials.ImapPassword = os.Getenv("IMAP_PASSWORD")

	// get port, then convert to uint32 data type
	port, err := strconv.ParseUint(os.Getenv("IMAP_PORT"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	// get tls, then convert to boolean data type
	tls, err := strconv.ParseBool(os.Getenv("IMAP_TLS"))
	if err != nil {
		log.Fatal(err)
	}

	cl.Credentials.ImapPort = uint32(port)
	cl.Credentials.ImapTLS = tls
}

// SetImapServer for set the imap server endpoint with format `hostname:port`
func (cl *Client) setImapServer() {
	cl.Credentials.ImapServer = fmt.Sprintf("%s:%d", cl.Credentials.ImapHostname, cl.Credentials.ImapPort)
}
