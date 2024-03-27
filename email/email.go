package email

import (
	"bytes"
	"errors"
	"text/template"

	"github.com/ibilalkayy/flow/db/budget_db"
	"github.com/ibilalkayy/flow/internal/middleware"
	"github.com/ibilalkayy/flow/internal/structs"
	"gopkg.in/gomail.v2"
)

func ViewAlert(category string) ([3]string, error) {
	ev := new(structs.EmailVariables)

	db, err := budget_db.Connection()
	if err != nil {
		return [3]string{}, err
	}

	query := "SELECT categories, category_amounts FROM Alert WHERE categories=$1"
	rows, err := db.Query(query, category)
	if err != nil {
		return [3]string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&ev.Category, &ev.CategoryAmount); err != nil {
			return [3]string{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return [3]string{}, err
	}

	values := [3]string{ev.Category, ev.CategoryAmount}
	return values, nil
}

func SendAlertEmail(category string) error {
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	myUsername := middleware.LoadEnvVariable("USERNAME")

	emailCreds, err := ViewAlert(category)
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
