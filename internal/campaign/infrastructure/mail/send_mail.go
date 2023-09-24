package infra_mail

import (
	campaign "marketplace/internal/campaign/domain"
	"os"

	"gopkg.in/gomail.v2"
)

func SendMail(campaign *campaign.Campaign) error {
	d := gomail.NewDialer(os.Getenv("SMTP"),587,os.Getenv("FROM_EMAIL"), os.Getenv("FROM_EMAIL_PASSWORD"))

	var emails []string
	for _,contac := range campaign.Contacts{
		emails = append(emails, contac.Email)
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", os.Getenv("FROM_EMAIL"))
	msg.SetHeader("To", emails...)//envia varios emails
	msg.SetHeader("Subject", campaign.Name)
	msg.SetBody("text/html",campaign.Content)

	return d.DialAndSend(msg)
}