package dbmodels

import (
	"time"

	"github.com/lib/pq"
)

type Notification struct {
	ID                  uint           `gorm:"primaryKey"`
	UserEmail           string         `gorm:"not null"`
	FromStationID       int64          `gorm:"not null"`
	ToStationID         int64          `gorm:"not null"`
	RouteID             string         `gorm:"not null"`
	SelectedSeatClasses pq.StringArray `gorm:"type:text[];not null"`
	NotificationTime    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	SeatClass           string         `gorm:"type:varchar(255)"`
	User                User           `gorm:"foreignKey:UserEmail;references:Email"`
}
