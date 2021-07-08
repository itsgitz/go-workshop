package imap

import "fmt"

// CheckNewEmail for checking the last message
// return true if a new email is arrived
func (cl *Client) CheckNewEmail() {
	cl.selectInbox()

	fmt.Println("inbox", cl.Mailbox.Status)
}
