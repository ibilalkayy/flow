package email

import (
	"bytes"
	"errors"
	"log"
	"text/template"

	"github.com/ibilalkayy/flow/entities"
	"github.com/ibilalkayy/flow/handler"
	"gopkg.in/gomail.v2"
)

type MyEmail struct {
	*handler.Handler
}

func (h MyEmail) SendAlertEmail(category string) error {
	myEmail := h.Deps.Env.LoadEnvVariable("APP_EMAIL")
	myPassword := h.Deps.Env.LoadEnvVariable("APP_PASSWORD")
	myUsername := h.Deps.Env.LoadEnvVariable("USERNAME")

	details, err := h.Deps.ManageBudget.ViewBudget(category)
	if err != nil {
		return err
	}

	mail := gomail.NewMessage()

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("framework/email/templates/alert.html")
	if err != nil {
		log.Fatal(err)
	}

	categoryName, ok1 := details[1].(string)
	categoryAmount, ok2 := details[2].(int)

	if !ok1 || !ok2 {
		return errors.New("unable to convert category and amount to int and string")
	}

	emailVariables := entities.EmailVariables{
		Username:       myUsername,
		Category:       categoryName,
		CategoryAmount: categoryAmount,
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
