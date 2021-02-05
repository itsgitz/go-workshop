package imap

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"github.com/emersion/go-imap"
	gi "github.com/emersion/go-imap"
	"github.com/emersion/go-message/mail"
)

// Mailbox data collection
type Mailbox struct {
	Status *gi.MailboxStatus
}

// ListMailBoxes show the list of mailboxes of imap client
func (cl *Client) ListMailBoxes() {
	done := make(chan error, 1)
	mailboxes := make(chan *gi.MailboxInfo, 10)

	// get list of mailboxes
	go func() {
		done <- cl.Client.List("", "*", mailboxes)
	}()

	for m := range mailboxes {
		log.Println("* Name", m.Name, ", Attributes[]", m.Attributes, ", Delimiter", m.Delimiter)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}
}

// selectInbox for select the INBOX mailbox
func (cl *Client) selectInbox() {
	inbox, err := cl.Client.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}

	cl.Mailbox.Status = inbox
}

// GetMessages for get all the last 4 messages
func (cl *Client) GetMessages() {
	// first, select the inbox mailbox
	cl.selectInbox()

	from := uint32(1)
	to := cl.Mailbox.Status.Messages

	if cl.Mailbox.Status.Messages > 3 {
		from = cl.Mailbox.Status.Messages - 1
	}

	log.Println("from:", from, "to:", to)

	seqset := new(gi.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *gi.Message, 10)
	done := make(chan error, 1)

	go func() {
		done <- cl.Client.Fetch(seqset, []gi.FetchItem{gi.FetchEnvelope}, messages)
	}()

	for m := range messages {
		log.Println("* Subject:", m.Envelope.Subject)
		log.Println("* Message ID:", m.Envelope.MessageId)

		for v, k := range m.Envelope.From {
			log.Println("From:", v, k)
		}

		for v, k := range m.Envelope.To {
			log.Println("To:", v, k)
		}
	}
}

// GetMessageBody for parse the e-mail body
func (cl *Client) GetMessageBody(wg *sync.WaitGroup) {
	cl.selectInbox()

	// check if the message is not empty
	if cl.Mailbox.Status.Messages == 0 {
		log.Fatal("No message in mailbox")
	}

	seqset := new(gi.SeqSet)
	seqset.AddNum(cl.Mailbox.Status.Messages)

	// get the whole message body
	var section gi.BodySectionName
	items := []imap.FetchItem{section.FetchItem()}

	messages := make(chan *gi.Message, 1)
	go func() {
		cl.Mutex.Lock()
		if err := cl.Client.Fetch(seqset, items, messages); err != nil {
			log.Fatal(err)
		}
		cl.Mutex.Unlock()

		wg.Done()
	}()

	m := <-messages
	if m == nil {
		log.Fatal("Server didn't returned message")
	}

	// get the email body
	r := m.GetBody(&section)
	if r == nil {
		log.Fatal("Server didn't returned body message")
	}

	// create a new mail reader
	reader, err := mail.CreateReader(r)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Header")

	header := reader.Header
	if date, err := header.Date(); err == nil {
		log.Println("Date:", date)
	}

	if from, err := header.AddressList("From"); err == nil {
		log.Println("From Name:", from[0].Name)
		log.Println("From Address:", from[0].Address)
	}

	if to, err := header.AddressList("To"); err == nil {
		log.Println("To Name:", to[0].Name)
		log.Println("To Address:", to[0].Address)
	}

	if subject, err := header.Subject(); err == nil {
		log.Println("Subject:", subject)
	}

	messageID, err := header.MessageID()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message-Id:", messageID)

	p, err := reader.NextPart()
	if err != nil {
		log.Fatal(err)
	}

	switch h := p.Header.(type) {
	case *mail.InlineHeader:
		body, err := ioutil.ReadAll(p.Body)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Got text: \n\n")
		fmt.Println(string(body))

	case *mail.AttachmentHeader:
		filename, err := h.Filename()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Got attachment: \n\n")
		fmt.Println(filename)
	}
}
