package repositories

import (
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type WatchedRouteRepository struct {
	DB *gorm.DB
}

func (r *WatchedRouteRepository) Create(watchedRoute *dbmodels.WatchedRoute) error {
	return r.DB.Create(watchedRoute).Error
}

func (r *WatchedRouteRepository) Delete(id uint) error {
	return r.DB.Delete(&dbmodels.WatchedRoute{}, id).Error
}

func (r *WatchedRouteRepository) GetAllForEmail(email string) ([]dbmodels.WatchedRoute, error) {
	var watchedRoutes []dbmodels.WatchedRoute
	err := r.DB.Where("user_email = ?", email).Find(&watchedRoutes).Error
	return watchedRoutes, err
}

func (r *WatchedRouteRepository) GetAll() ([]dbmodels.WatchedRoute, error) {
	var watchedRoutes []dbmodels.WatchedRoute
	err := r.DB.Find(&watchedRoutes).Error
	return watchedRoutes, err
}
