package mailer

import (
	"net/mail"
	"net/smtp"
)

type MailerAddr = mail.Address

type MailerSMTP struct {
	SMTPHostname   string
	SMTPServername string
}

type MailerAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MailerDst struct {
	Name    string
	Message string
	Subject string
}

type Mail struct {
	To      string `json:"to" db:"to"`
	Subject string `json:"subject" db:"subject"`
	Message string `json:"message" db:"message"`
	DstName string `json:"destinationName" db:"destinationName"`
	Html    string `json:"html" db:"html"`
}

type Mailer struct {
	DstName  string
	FromAddr *mail.Address
	ToAddr   *mail.Address
	Subject  string
	Message  string
	Provider MailerSMTP
	Auth     MailerAuth
	Html     string
}

type MailerStructured struct {
	Mailer *Mailer
	Client *smtp.Client
}
