package dbmodels

import (
	"github.com/lib/pq"
	"time"
)

type SuccessfulPurchase struct {
	ID                  uint           `gorm:"primaryKey"`
	UserEmail           string         `gorm:"not null"`
	FromStationID       int64          `gorm:"not null"`
	ToStationID         int64          `gorm:"not null"`
	RouteID             string         `gorm:"not null"`
	TariffClass         string         `gorm:"not null"`
	SelectedSeatClasses pq.StringArray `gorm:"type:text[];not null"`
	PurchaseTime        time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	Price               float32        `gorm:"type:decimal(10,2)"`
	Currency            string         `gorm:"type:varchar(16)"`
	SeatClass           string         `gorm:"type:varchar(255)"`
	User                User           `gorm:"foreignKey:UserEmail;references:Email"`
}
