package email

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/ibilalkayy/flow/db/alert_db"
	"github.com/ibilalkayy/flow/internal/middleware"
	"github.com/ibilalkayy/flow/internal/structs"
	"gopkg.in/gomail.v2"
)

func SendAlertEmail(category string) error {
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	myUsername := middleware.LoadEnvVariable("USERNAME")

	emailCreds, err := alert_db.ViewAlert(category)
	if err != nil {
		return err
	}

	mail := gomail.NewMessage()

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("email/templates/alert.html")
	if err != nil {
		return err
	}

	emailVariables := structs.EmailVariables{
		Username:       myUsername,
		Category:       emailCreds[0],
		CategoryAmount: emailCreds[1],
	}

	if err := temp.Execute(body, emailVariables); err != nil {
		return errors.New("cannot load the template")
	}

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", myEmail)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", "Notification of your spending on flow")
	mail.SetBody("text/html", body.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := dialer.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
