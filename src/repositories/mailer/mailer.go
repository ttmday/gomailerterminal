package mailer

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/ttmday/gomailerterminal/src/global/constants/envs"
	tmpls "github.com/ttmday/gomailerterminal/src/templates"
)

var message = ""

func New(mail Mail, c *MailerAuth) *Mailer {
	return &Mailer{
		DstName: mail.DstName,
		FromAddr: &MailerAddr{
			Name:    "",
			Address: c.Username,
		},
		ToAddr: &MailerAddr{
			Name:    mail.DstName,
			Address: mail.To,
		},
		Subject: mail.Subject,
		Message: mail.Message,
		Provider: MailerSMTP{
			SMTPHostname:   envs.SMTP_Hostname,
			SMTPServername: envs.SMTP_Servername,
		},
		Auth: MailerAuth{
			Username: c.Username,
			Password: c.Password,
		},
	}
}

func (m *Mailer) CreateMail() (*smtp.Client, error) {
	from := m.FromAddr
	to := m.ToAddr
	subject := m.Subject

	dst := &MailerDst{Name: to.Name, Message: m.Message, Subject: m.Subject}

	headers := setHeaders(from.String(), to.String(), subject)

	message = generateMessageByHeaders(headers)

	t := &template.Template{}

	t = tmpls.LoadTemplate("src/templates/mail.html")

	message += tmpls.RenderTemplateByBuf(t, dst)

	provider := getProvider(m.Provider)

	m.Provider = *provider

	auth := generateAuth(m.Auth, provider.SMTPHostname)

	client, err := setConnection(provider, auth)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *Mailer) SendMail(client *smtp.Client) error {
	if err := client.Mail(m.FromAddr.Address); err != nil {
		println("Error sending mail Address from")
		return err
	}

	if err := client.Rcpt(m.ToAddr.Address); err != nil {
		println("Error sending mail Address to")
		return err
	}

	w, err := client.Data()
	if err != nil {
		println("Error sending mail")
		return err
	}

	_, err = w.Write([]byte(message))

	if err := w.Close(); err != nil {
		println("Error sending mail close message")
		return err
	}

	if err := client.Close(); err != nil {
		println("Error sending mail close client")
		return err
	}

	println("Mail sent successfully")
	return nil
}

func setHeaders(from, to, subject string) map[string]string {
	headers := make(map[string]string)

	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject
	headers["content-type"] = "text/html; charset=utf-8"

	return headers
}

func generateMessageByHeaders(headers map[string]string) string {
	message := ""

	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	return message
}

func getProvider(provider MailerSMTP) *MailerSMTP {
	p := &MailerSMTP{
		SMTPHostname:   "",
		SMTPServername: "",
	}

	if provider.SMTPHostname == "" {
		p.SMTPHostname = envs.SMTP_Hostname
	} else {
		p.SMTPHostname = provider.SMTPHostname
	}

	if provider.SMTPServername == "" {
		p.SMTPServername = envs.SMTP_Servername
	} else {
		p.SMTPServername = provider.SMTPServername
	}

	return p
}

func generateAuth(auth MailerAuth, host string) smtp.Auth {
	return smtp.PlainAuth("", auth.Username, auth.Password, host)
}

func setConnection(provider *MailerSMTP, auth smtp.Auth) (*smtp.Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         provider.SMTPHostname,
	}

	conn, err := tls.Dial("tcp", provider.SMTPServername, tlsConfig)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, tlsConfig.ServerName)

	client.Auth(auth)

	return client, nil
}
