package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/spf13/viper"
)

const (
	smtp_server = "smtp.gmail.com"
	smtp_email  = "kiettranuit@gmail.com"
	smtp_port   = 587

	verify_signup = "verify sign up"
)

var (
	smtp_url = fmt.Sprintf("%s:%d", smtp_server, smtp_port)
	mime     = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type MailService struct {
	Password string `mapstructure:"MAIL_PASSWORD"`
	Auth     smtp.Auth
}

func NewMailService(path string) (mail MailService, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("mail")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&mail)
	mail.Auth = smtp.PlainAuth("", smtp_email, mail.Password, smtp_server)

	return
}

type MailMessage struct {
	From    string
	To      []string
	Subject string
	Content string
}

func CreateVerificationMail(to []string, code string) MailMessage {
	return MailMessage{
		From:    smtp_email,
		To:      to,
		Subject: "Verify Email",
		Content: code,
	}
}

func (mail_service MailService) SendMail(mail MailMessage) error {
	buffer := new(bytes.Buffer)
	if tmpl, err := template.ParseFiles("website/verify.html"); err != nil {
		return err
	} else {
		tmpl.Execute(buffer, mail)
	}

	to := fmt.Sprintf("To: %s\r\n", mail.To[0])
	body := buffer.String()
	subject := "Subject: " + mail.Subject + "\n"
	message := to + subject + mime + "\r\n" + body

	return smtp.SendMail(smtp_url, mail_service.Auth, mail.From, mail.To, []byte(message))
}
