package repositories

import (
	"fmt"

	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type SuccessfulPurchaseRepository struct {
	DB *gorm.DB
}

func (r *SuccessfulPurchaseRepository) Create(successfulPurchase *dbmodels.SuccessfulPurchase) error {
	return r.DB.Create(successfulPurchase).Error
}

func (r *SuccessfulPurchaseRepository) CreateWithBeerUpdate(
	successfulPurchase *dbmodels.SuccessfulPurchase, beerIncrement float32,
) (float32, error) {
	var newBeersOwed float32

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(successfulPurchase).Error; err != nil {
			return err
		}

		if err := tx.Model(&dbmodels.User{}).Where("email = ?", successfulPurchase.UserEmail).
			Update("beers_owed", gorm.Expr("beers_owed + ?", beerIncrement)).Error; err != nil {
			return err
		}

		if err := tx.Model(&dbmodels.User{}).Where("email = ?", successfulPurchase.UserEmail).
			Select("beers_owed").Scan(&newBeersOwed).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("failed to create successful purchase with beer update: %w", err)
	}

	return newBeersOwed, nil
}
