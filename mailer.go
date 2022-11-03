package main

import (
	"errors"
	"fmt"
	"net/smtp"
)

type Mailer struct {
	auth           smtp.Auth
	smtpHost       string
	smtpPort       uint16
	smtpAddr       string
	smtpAuthUser   string
	smtpAuthPass   *string
	smtpSenderUser string
}

func NewMailer(
	smtpHost string,
	smtpPort uint16,
	smtpAuthUser string,
	smtpAuthPass *string,
	smtpSenderUser string,
) *Mailer {
	mailer := &Mailer{
		smtpHost:       smtpHost,
		smtpPort:       smtpPort,
		smtpAuthUser:   smtpAuthUser,
		smtpAuthPass:   smtpAuthPass,
		smtpSenderUser: smtpSenderUser,
	}

	mailer.smtpAddr = fmt.Sprintf("%s:%d", mailer.smtpHost, mailer.smtpPort)

	var password string
	if mailer.smtpAuthPass != nil {
		password = *mailer.smtpAuthPass
	}

	mailer.auth = smtp.PlainAuth("", mailer.smtpAuthUser, password, mailer.smtpHost)

	return mailer
}

func (m *Mailer) Send(
	senderName *string,
	destination *string,
	subject *string,
	bodyMessage string,
) error {
	if senderName == nil {
		return errors.New("senderName is required")
	}

	if destination == nil {
		return errors.New("destination is required")
	}

	if subject == nil {
		return errors.New("subject is required")
	}

	body := "From: " + *senderName + " <" + m.smtpSenderUser + ">\n" +
		"Subject: " + *subject + "\n\n" +
		bodyMessage

	if err := smtp.SendMail(m.smtpAddr, m.auth, m.smtpSenderUser, []string{*destination}, []byte(body)); err != nil {
		return err
	}

	return nil
}
