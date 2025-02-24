package repositories

import (
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func (r *NotificationRepository) Create(notification *dbmodels.Notification) error {
	return r.DB.Create(notification).Error
}
