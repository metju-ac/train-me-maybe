package notification

import (
	"fmt"
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	gomail "gopkg.in/mail.v2"
)

func formatConnectionShort(input *models.UserInput) string {
	return fmt.Sprintf("%s -> %s on %s", input.DepartingStation.City, input.ArrivingStation.City, input.SelectedRoute.DepartureTime.Format("02.01. 15:04"))
}

func sendEmail(config *config.SmtpConfig, subject string, body string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Username)

	// Set E-Mail receivers
	m.SetHeader("To", config.Recipient)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", body)

	// Settings for SMTP server
	d := gomail.NewDialer(config.Server, config.Port, config.Username, config.Password)

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		slog.Error("Failed to send email", "error", err)
		panic(err)
	}

	slog.Info("Email notification sent successfully")
	return
}

// partly taken from https://www.loginradius.com/blog/engineering/sending-emails-with-golang/
func EmailNotificationFreeSeats(config *config.SmtpConfig, input *models.UserInput) {
	slog.Info("Preparing email notification", "departingStation", input.DepartingStation.StationID, "arrivingStation", input.ArrivingStation.StationID, "selectedRoutes", input.SelectedRoute.Id)

	subject := `[REGIOJET] Free seats found on ` + formatConnectionShort(input)
	body := `There are free seats on the route from <b>` + cli.FormatStation(input.DepartingStation) + `</b> to <b>` + cli.FormatStation(input.ArrivingStation) + `</b> on the selected route: ` + cli.FormatRoute(input.SelectedRoute)
	sendEmail(config, subject, body)

	return
}

func EmailNotificationTicketBought(config *config.SmtpConfig, input *models.UserInput) {
	slog.Info("Preparing email notification", "departingStation", input.DepartingStation.StationID, "arrivingStation", input.ArrivingStation.StationID, "selectedRoutes", input.SelectedRoute.Id)

	subject := `[REGIOJET] Ticket purchase successful for ` + formatConnectionShort(input)
	body := `Your ticket has been successfully purchased for the route from <b>` + cli.FormatStation(input.DepartingStation) + `</b> to <b>` + cli.FormatStation(input.ArrivingStation) + `</b> on the selected route: ` + cli.FormatRoute(input.SelectedRoute)
	sendEmail(config, subject, body)

	return
}
