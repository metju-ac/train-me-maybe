package repositories

import (
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *dbmodels.User) error
	FindByEmail(email string) (*dbmodels.User, error)
	Update(user *dbmodels.User) error
	DeleteByEmail(email string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *dbmodels.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*dbmodels.User, error) {
	var user dbmodels.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *dbmodels.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteByEmail(email string) error {
	return r.db.Where("email = ?", email).Delete(&dbmodels.User{}).Error
}
