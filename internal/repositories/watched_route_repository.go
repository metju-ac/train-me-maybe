package repositories

import (
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type WatchedRouteRepository interface {
	Create(watchedRoute *dbmodels.WatchedRoute) error
	Delete(id uint) error
	GetAllForEmail(email string) ([]dbmodels.WatchedRoute, error)
	GetAll() ([]dbmodels.WatchedRoute, error)
}

type watchedRouteRepository struct {
	db *gorm.DB
}

func NewWatchedRouteRepository(db *gorm.DB) WatchedRouteRepository {
	return &watchedRouteRepository{db: db}
}

func (r *watchedRouteRepository) Create(watchedRoute *dbmodels.WatchedRoute) error {
	return r.db.Create(watchedRoute).Error
}

func (r *watchedRouteRepository) Delete(id uint) error {
	return r.db.Delete(&dbmodels.WatchedRoute{}, id).Error
}

func (r *watchedRouteRepository) GetAllForEmail(email string) ([]dbmodels.WatchedRoute, error) {
	var watchedRoutes []dbmodels.WatchedRoute
	err := r.db.Where("user_email = ?", email).Find(&watchedRoutes).Error
	return watchedRoutes, err
}

func (r *watchedRouteRepository) GetAll() ([]dbmodels.WatchedRoute, error) {
	var watchedRoutes []dbmodels.WatchedRoute
	err := r.db.Find(&watchedRoutes).Error
	return watchedRoutes, err
}
