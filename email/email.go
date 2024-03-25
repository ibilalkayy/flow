package email

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"text/template"

	internal_alert "github.com/ibilalkayy/flow/internal/app/alert"
	"github.com/ibilalkayy/flow/internal/middleware"
	"github.com/ibilalkayy/flow/internal/structs"
	"gopkg.in/gomail.v2"
)

func SendAlertMail() {
	var myCategoryAmount, myTotalAmount string
	// var ev structs.EmailVariables
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")
	myUsername := middleware.LoadEnvVariable("USERNAME")
	// myCategory := ev.Category

	emailCreds, err := internal_alert.ViewEmailCredentials("total_category_w4csyvdm")
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(emailCreds[0], "total_category") {
		_ = strings.TrimPrefix(emailCreds[0], "total_category")
		myTotalAmount = emailCreds[2]
	} else {
		myCategoryAmount = emailCreds[1]
	}

	mail := gomail.NewMessage()

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("email/templates/alert.html")
	if err != nil {
		log.Fatal(err)
	}

	emailVariables := structs.EmailVariables{
		Username:       myUsername,
		Category:       emailCreds[0],
		CategoryAmount: myCategoryAmount,
		TotalAmount:    myTotalAmount,
	}

	if err := temp.Execute(body, emailVariables); err != nil {
		fmt.Println(errors.New("cannot load the template"))
	}

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", myEmail)
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", "Notification of your spending on flow")
	mail.SetBody("text/html", body.String())

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}
