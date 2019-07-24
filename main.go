package main

import (
	"context"
	"log"
	"net/smtp"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var weltsparenEmail = os.Getenv("WELTSPAREN_EMAIL")
	var weltsparenPassword = os.Getenv("WELTSPAREN_PASSWORD")
	var notificationEmailFrom = os.Getenv("NOTIFICATION_EMAIL_FROM")
	var notificationEmailTo = os.Getenv("NOTIFICATION_EMAIL_TO")
	var notificationEmailPassword = os.Getenv("NOTIFICATION_EMAIL_PASSWORD")
	var notificationEmailServer = os.Getenv("NOTIFICATION_EMAIL_SERVER")
	var notificationEmailPort = os.Getenv("NOTIFICATION_EMAIL_PORT")

	value, err := weltsparen(weltsparenEmail, weltsparenPassword)
	if err != nil {
		log.Fatal(err)
	}

	err = notify(
		notificationEmailServer,
		notificationEmailPort,
		notificationEmailFrom,
		notificationEmailTo,
		notificationEmailPassword,
		value,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func weltsparen(email string, password string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithErrorf(log.Printf))
	defer cancel()

	var res string
	err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate(`https://banking.weltsparen.de/savingglobal/part/Welcome/content/loginEmail`),
			chromedp.WaitVisible(`input[name="j_username"]`),
			chromedp.SetValue(`input[name="j_username"]`, email),
			chromedp.SetValue(`input[name="j_password"]`, password),
			chromedp.Click(`.evt-login`),
			chromedp.WaitVisible(`#mainDonutInnerValue`),
			chromedp.Text(`#mainDonutInnerValue`, &res),
		},
	)

	return res, err
}

func notify(
	server string,
	port string,
	from string,
	to string,
	password string,
	value string,
) error {
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Weltsparen Update\n\n" +
		"Investments Total Value: € " + value + "\n\n"

	return smtp.SendMail(
		(server + ":" + port),
		smtp.PlainAuth("", from, password, server),
		from,
		[]string{to},
		[]byte(msg),
	)
}
