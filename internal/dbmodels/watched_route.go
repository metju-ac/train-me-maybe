package dbmodels

import "github.com/lib/pq"

type WatchedRoute struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	UserEmail           string         `gorm:"not null" json:"userEmail"`
	FromStationID       int64          `gorm:"not null" json:"fromStationId"`
	ToStationID         int64          `gorm:"not null" json:"toStationId"`
	RouteID             string         `gorm:"not null" json:"routeId"`
	SelectedSeatClasses pq.StringArray `gorm:"type:text[];not null" json:"selectedSeatClasses"`
	AutoPurchase        bool           `gorm:"not null;default:false" json:"autoPurchase"`
	TariffClass         string         `json:"tariffClass"`
	CreditUserNumber    string         `json:"creditUser"`
	CreditUserPassword  string         `json:"creditPassword"`
	User                User           `gorm:"foreignKey:UserEmail;references:Email" json:"-"`
}
