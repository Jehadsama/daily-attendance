package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/Jehadsama/daily-attendance/internal/utils"
)

var (
	smtpHost     string = os.Getenv("smtpHost")
	smtpPort     string = os.Getenv("smtpPort")
	smtpFrom     string = os.Getenv("smtpFrom")
	smtpPassword string = utils.ReadFile("smtpPassword")

	auth = smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)
)

type EmailData struct {
	From string
	To   []string

	Subject string
	Content string
}

func (ed EmailData) Format() []byte {
	return []byte(
		"From: " + ed.From + "\r\n" +
			"To: " + ed.To[0] + "\r\n" +
			"Subject: " + ed.Subject + "\r\n" +
			"\r\n" +
			ed.Content + "\r\n")
}

func (ed EmailData) SendMail() error {
	err := smtp.SendMail(smtpHost+smtpPort, auth, smtpFrom, ed.To, ed.Format())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
