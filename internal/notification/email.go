package notification

import (
	"fmt"
	openapiclient "github.com/metju-ac/train-me-maybe/openapi"
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/cli"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"github.com/metju-ac/train-me-maybe/internal/models"
	gomail "gopkg.in/mail.v2"
)

func formatConnectionShort(input *models.UserInput) string {
	return fmt.Sprintf("%s -> %s on %s", input.DepartingStation.City, input.ArrivingStation.City, input.RouteDetail.DepartureTime.Format("02.01. 15:04"))
}

func formatRoute(route *openapiclient.Route) string {
	return fmt.Sprintf("%s - %s (Transfers: %d)", route.DepartureTime.Format("02.01. 15:04"), route.ArrivalTime.Format("02.01. 15:04"), len(route.GetTransfersInfo().Transfers))
}

func sendEmail(config *config.SmtpConfig, input *models.UserInput, subject string, body string) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.Username)

	// Set E-Mail receivers
	m.SetHeader("To", input.UserEmail)

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
	slog.Info("Preparing free seats email notification", "departingStation", input.DepartingStation.StationID, "arrivingStation", input.ArrivingStation.StationID, "selectedRoutes", input.SelectedRouteIds)

	subject := `[TRAIN ME MAYBE] Free seats found on ` + formatConnectionShort(input)
	body := `There are free seats on the route from <b>` + cli.FormatStation(input.DepartingStation) + `</b> to <b>` + cli.FormatStation(input.ArrivingStation) + `</b> on the selected route: ` + formatRoute(input.RouteDetail)
	sendEmail(config, input, subject, body)

	return
}

func generateTicketBoughtEmailBody(input *models.UserInput, ticket *openapiclient.Ticket, beersOwed *float32) string {
	body := `Your ticket has been successfully purchased for the route from <b>` +
		cli.FormatStation(input.DepartingStation) + `</b> to <b>` + cli.FormatStation(input.ArrivingStation) +
		`</b> on the selected route: ` + formatRoute(input.RouteDetail) + ` for <b>` +
		fmt.Sprintf("%.2f", ticket.Price) + ` ` + string(ticket.Currency) + `</b>.`
	if beersOwed != nil {
		body += `<br><br> You now owe us ` + fmt.Sprintf("%.2f", *beersOwed) + ` beers for our amazing service. Thanks for keeping Robert and Matěj hydrated 🍻.`
	}
	return body
}

func EmailNotificationTicketBought(config *config.SmtpConfig, input *models.UserInput, ticket *openapiclient.Ticket) {
	slog.Info("Preparing bought ticket email notification", "departingStation", input.DepartingStation.StationID, "arrivingStation", input.ArrivingStation.StationID, "selectedRoutes", input.SelectedRouteIds)

	subject := `[TRAIN ME MAYBE] Ticket purchase successful for ` + formatConnectionShort(input)
	body := generateTicketBoughtEmailBody(input, ticket, nil)
	sendEmail(config, input, subject, body)
}

func EmailNotificationTicketBoughtWithBeers(config *config.SmtpConfig, input *models.UserInput, ticket *openapiclient.Ticket, beersOwed float32) {
	slog.Info("Preparing bought ticket email notification with beers owed", "departingStation", input.DepartingStation.StationID, "arrivingStation", input.ArrivingStation.StationID, "selectedRoutes", input.SelectedRouteIds)

	subject := `[TRAIN ME MAYBE] Ticket purchase successful for ` + formatConnectionShort(input)
	body := generateTicketBoughtEmailBody(input, ticket, &beersOwed)
	sendEmail(config, input, subject, body)
}

func EmailNotificationTicketNotPaid(config *config.SmtpConfig, input *models.UserInput, price float32, currency string) {
	slog.Info("Preparing not paid ticket email notification", "price", price, "currency", currency)

	subject := `[TRAIN ME MAYBE] Ticket purchase failed`
	body := `Your ticket purchase has failed. The ticket for <b>` + fmt.Sprintf("%.2f", price) + ` ` + currency + `</b> has been booked but not been paid. Please, check that you have enough credit. If you do, please report this issue to the support.`
	sendEmail(config, input, subject, body)

	return
}

func EmailNotificationLowCredit(config *config.SmtpConfig, input *models.UserInput, remainingCredit float64, currency string) {
	slog.Info("Preparing low credit email notification", "remainingCredit", remainingCredit, "currency", currency)

	subject := `[TRAIN ME MAYBE] Low credit alert`
	body := `Your credit is running low. Your remaining credit is <b>` + fmt.Sprintf("%.2f", remainingCredit) + ` ` + currency + `</b>.`
	sendEmail(config, input, subject, body)

	return
}
