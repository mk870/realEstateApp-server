package services

import (
	"net/smtp"
)

func emailHtml(firstName string, frontendLink string) string {
	return ("<h1>Hi " + firstName + "</h1> <p>Thank you for registering with <span style=\"color: rgb(71,91,232);font-size: 18px;font-weight:bold;\">R Estates</span>. Please click on the link to activate your account: <a href=\"" + frontendLink + "\"style=\"font-size: 15px;\">Activate Now</a></p></p> Link will expire in 15 minutes. See you soon</p>")
}

func SendVerificationEmail(email string, firstName string, frontendLink string) bool {
	html := emailHtml(firstName, frontendLink)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: Confirm your email for Account Registration\n" + headers + "\n\n" + html
	auth := smtp.PlainAuth(
		"",
		"webdevndlovu5@gmail.com",
		"tdjwgzmxbhmgydjl",
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"webdevndlovu5@gmail.com",
		[]string{email},
		[]byte(msg),
	)
	if err == nil {
		return true
	} else {
		return false
	}
}
