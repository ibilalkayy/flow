package email

import (
	"log"

	"github.com/ibilalkayy/flow/internal/middleware"
	"gopkg.in/gomail.v2"
)

func SendAlertMail(value [2]string) {
	mail := gomail.NewMessage()

	// body := new(bytes.Buffer)
	// temp, err := template.ParseFiles("/templates/" + value[0] + ".html")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")

	// if err := temp.Execute(body, ); err != nil {
	// 	fmt.Println(errors.New("Cannot load the template"))
	// }

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", value[0])
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", value[1])
	mail.SetBody("text/html", "<h1> Hello there, how are you doing today? </h1>")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}
