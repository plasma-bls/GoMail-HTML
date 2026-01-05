package main

import (
	"encoding/json"
	"net/smtp"
	"os"
	"strings"
)

type Data struct {
	Smtphost string `json:"smtphost"`
	Smtport  string `json:"smtport"`
	Password string `json:"password"`
	Username string `json:"username"`
	Emails   string `json:"emails"`
	Subject  string `json:"subject"`
	Mime     string `json:"mime"`
	Html     string `json:"html"`
}

func main() {

	// extracting values from the "config.json" file
	content, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var payload Data
	if err = json.Unmarshal(content, &payload); err != nil {
		panic(err)
	}

	// extracting content from the HTML file that we will later use for our email.
	htcontent, err := os.ReadFile(payload.Html)

	// authentication to the sender account
	auth := smtp.PlainAuth("", payload.Username, payload.Password, payload.Smtphost)

	// preparing the recipient list (if there are more than 1 receivers)
	recipients := strings.Split(payload.Emails, ",")
	for i := range recipients {
		recipients[i] = strings.TrimSpace(recipients[i])
	}

	// constructing the email message
	msg := payload.Subject + payload.Mime + string(htcontent)

	// sending emails
	if err = smtp.SendMail(payload.Smtphost+":"+payload.Smtport, auth, payload.Username, recipients, []byte(msg)); err != nil {
		panic(err)
	}
}
