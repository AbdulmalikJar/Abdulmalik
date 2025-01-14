package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"gopkg.in/gomail.v2"
)

func sendScheduledEmail() {
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")
	recipient := os.Getenv("RECIPIENT_EMAIL")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", email)
	mailer.SetHeader("To", recipient)
	mailer.SetHeader("Subject", "Automated Scheduled Email")
	mailer.SetBody("text/plain", "Hello! This is an automated email sent every 12 hours via our Go email scheduler!")

	port := 587
	if smtpPort != "" {
		_, err := fmt.Sscanf(smtpPort, "%d", &port)
		if err != nil {
			return
		}
	}
	dialer := gomail.NewDialer(smtpServer, port, email, password)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("Failed to send email: %v\n", err)
	} else {
		log.Println("Email sent successfully!")
	}
}

func main() {
	c := cron.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, err = c.AddFunc("@every 12h", func() {
		log.Println("Running scheduled email job...")
		sendScheduledEmail()
	})
	if err != nil {
		log.Fatalf("Failed to schedule email job: %v\n", err)
	}

	c.Start()

	log.Println("Email scheduler is running. Press Ctrl+C to stop.")
	select {}
}
