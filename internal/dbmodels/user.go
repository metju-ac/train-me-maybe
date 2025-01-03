package dbmodels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email              string `gorm:"uniqueIndex;not null"`
	PasswordHash       string `gorm:"not null"`
	Salt               string `gorm:"not null"`
	CreditUserNumber   *string
	CreditUserPassword *string
	CutOffTime         *int
	MinimalCredit      *int
	DefaultTariffClass *string
	BeersOwed          float64 `gorm:"not null;default:0"`
}
