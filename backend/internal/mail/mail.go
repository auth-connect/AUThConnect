package mail

import (
	"bytes"
	"crypto/tls"
	"embed"
	"text/template"
	"time"

	"github.com/go-mail/mail/v2"
	"github.com/google/uuid"
)

//go:embed "templates"
var templateFS embed.FS

type Mail struct {
	dialer *mail.Dialer
	sender string
}

func New(host string, port int, username, password, sender string) Mail {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = 5 * time.Second

	dialer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	return Mail{
		dialer: dialer,
		sender: sender,
	}
}

func (m Mail) Send(recipient, templateFile string, payload any) error {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", payload)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", payload)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", payload)
	if err != nil {
		return err
	}

	messageID := "<" + uuid.New().String() + "@" + "auth-connect.gr" + ">" // Use your domain here

	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetHeader("Message-ID", messageID)
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	for i := 1; i <= 3; i++ {
		err = m.dialer.DialAndSend(msg)
		if nil == err {
			return nil
		}

		time.Sleep(500 * time.Millisecond)
	}

	return err
}
