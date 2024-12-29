package purchase

import (
	"log/slog"

	"github.com/metju-ac/train-me-maybe/internal/config"
)

func AutoPurchaseTicket(config *config.Config) error {
	if !config.Auth.CreditEnabled {
		return nil
	}

	slog.Info("Auto purchasing ticket")

	return nil
}
