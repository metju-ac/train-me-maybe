package notification

import (
	"fmt"
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/handlers"
	gomail "gopkg.in/mail.v2"
)

func formatConnectionShort(route *handlers.HandleRouteSelectionResponse) string {
	return fmt.Sprintf("%s -> %s on %s", route.DepartingStation.City, route.ArrivingStation.City, route.SelectedRoute.DepartureTime.Format("02.01. 15:04"))
}

// partly taken from https://www.loginradius.com/blog/engineering/sending-emails-with-golang/
func EmailNotification(config *config.SmtpConfig, route *handlers.HandleRouteSelectionResponse) {
	slog.Info("Preparing email notification", "departingStation", route.DepartingStation.StationID, "arrivingStation", route.ArrivingStation.StationID, "selectedRoutes", route.SelectedRoute.Id)

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Username)

	// Set E-Mail receivers
	m.SetHeader("To", config.Recipient)

	// Set E-Mail subject
	m.SetHeader("Subject", `[REGIOJET] free seats found on `+formatConnectionShort(route))

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", `There are free seats on the route from <b>`+cli.FormatStation(route.DepartingStation)+`</b> to <b>`+cli.FormatStation(route.ArrivingStation)+`</b> on the selected route: `+cli.FormatRoute(route.SelectedRoute))

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
