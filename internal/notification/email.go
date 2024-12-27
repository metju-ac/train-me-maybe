package notification

import (
	"fmt"

	"github.com/metju-ac/train-me-maybe/internal/config"
	gomail "gopkg.in/mail.v2"
)

// partly taken from https://www.loginradius.com/blog/engineering/sending-emails-with-golang/
func EmailNotification(config *config.SmtpConfig, departingStation, arrivingStation int64, selectedRoutes []int64, seatClasses []string) {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Username)

	// Set E-Mail receivers
	m.SetHeader("To", config.Recipient)

	// Set E-Mail subject
	m.SetHeader("Subject", "[REGIOJET] free seats found")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", `There are free seats on the route from `+fmt.Sprint(departingStation)+` to `+fmt.Sprint(arrivingStation)+` on the following routes: `+fmt.Sprint(selectedRoutes)+` in the following seat classes: `+fmt.Sprint(seatClasses))

	// Settings for SMTP server
	d := gomail.NewDialer(config.Server, config.Port, config.Username, config.Password)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
