package dbmodels

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email              string  `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash       string  `gorm:"not null"             json:"-"`
	Salt               string  `gorm:"not null"             json:"-"`
	CreditUserNumber   *string `json:"creditUser"`
	CreditUserPassword *string `json:"creditPassword"`
	CutOffTime         *int    `json:"cutOffTime"`
	MinimalCredit      *int    `json:"minimalCredit"`
	DefaultTariffClass *string `json:"tariffKey"`
	BeersOwed          float64 `gorm:"not null;default:0"   json:"debt"`
}
