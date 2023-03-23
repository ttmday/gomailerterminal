package mailer

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	consts "github.com/ttmday/gomailerterminal/src/global/constants"
	tmpls "github.com/ttmday/gomailerterminal/src/template"
	"github.com/ttmday/gomailerterminal/src/views"
)

var message = ""

func Usage() {
	fmt.Println(" **********************************************************************")
	fmt.Println(" *                                                                    *")
	fmt.Println(" *        1) gomailer -u email -p password                            *")
	fmt.Println(" *        2) gomailer -f /path/to/credentials.json                    *")
	fmt.Println(" *                                                                    *")
	fmt.Println(" *   {'username': 'mail@example.com', 'password': 'password'}         *")
	fmt.Println(" *                                                                    *")
	fmt.Println(" **********************************************************************")
}

func New(mail *Mail, c *MailerAuth) *Mailer {
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
			SMTPHostname:   consts.SMTP_Hostname,
			SMTPServername: consts.SMTP_Servername,
		},
		Auth: MailerAuth{
			Username: c.Username,
			Password: c.Password,
		},
	}
}

func (m *Mailer) CreateMail() (*MailerStructured, error) {
	from := m.FromAddr
	to := m.ToAddr
	subject := m.Subject

	dst := &MailerDst{Name: to.Name, Message: m.Message, Subject: m.Subject}

	headers := setHeaders(from.String(), to.String(), subject)

	message = generateMessageByHeaders(headers)

	t := &template.Template{}

	t = tmpls.ParseView(views.MailView)

	message += tmpls.RenderTemplateByBuf(t, dst)

	provider := getProvider()

	m.Provider = *provider

	auth := generateAuth(m.Auth, provider.SMTPHostname)

	client, err := setConnection(provider, auth)
	if err != nil {
		return nil, err
	}

	return &MailerStructured{
		Client: client,
		Mailer: m,
	}, nil
}

func (m *MailerStructured) SendMail() (bool, error) {
	if err := m.Client.Mail(m.Mailer.FromAddr.Address); err != nil {
		println("Error sending mail Address from")
		return false, err
	}

	if err := m.Client.Rcpt(m.Mailer.ToAddr.Address); err != nil {
		println("Error sending mail Address to")
		return false, err
	}

	w, err := m.Client.Data()
	if err != nil {
		println("Error sending mail")
		return false, err
	}

	_, err = w.Write([]byte(message))

	if err := w.Close(); err != nil {
		println("Error sending mail close message")
		return false, err
	}

	if err := m.Client.Close(); err != nil {
		println("Error sending mail close client")
		return false, err
	}

	return true, nil
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

func getProvider() *MailerSMTP {
	godotenv.Load(".env")

	SMTP_Hostname := os.Getenv("SMTP_Hostname")
	SMTP_Servername := os.Getenv("SMTP_Servername")

	if SMTP_Hostname == "" {
		SMTP_Hostname = consts.SMTP_Hostname
	}

	if SMTP_Servername == "" {
		SMTP_Servername = consts.SMTP_Servername
	}

	return &MailerSMTP{
		SMTPHostname:   SMTP_Hostname,
		SMTPServername: SMTP_Servername,
	}
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

	if err != nil {
		return nil, err
	}

	if err = client.Auth(auth); err != nil {
		return nil, err
	}

	return client, nil
}
