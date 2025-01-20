package repositories

import (
	"fmt"

	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type SuccessfulPurchaseRepository interface {
	Create(successfulPurchase *dbmodels.SuccessfulPurchase) error
	CreateWithBeerUpdate(successfulPurchase *dbmodels.SuccessfulPurchase, beerIncrement float32) (float32, error)
}

type successfulPurchaseRepository struct {
	db *gorm.DB
}

func NewSuccessfulPurchaseRepository(db *gorm.DB) SuccessfulPurchaseRepository {
	return &successfulPurchaseRepository{db: db}
}

func (r *successfulPurchaseRepository) Create(successfulPurchase *dbmodels.SuccessfulPurchase) error {
	return r.db.Create(successfulPurchase).Error
}

func (r *successfulPurchaseRepository) CreateWithBeerUpdate(
	successfulPurchase *dbmodels.SuccessfulPurchase, beerIncrement float32,
) (float32, error) {
	var newBeersOwed float32

	err := r.db.Transaction(func(tx *gorm.DB) error {
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
