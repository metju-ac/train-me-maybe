package repositories

import (
	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *dbmodels.User) error
	FindByEmail(email string) (*dbmodels.User, error)
	Update(user *dbmodels.User) error
	DeleteByEmail(email string) error
}

type userRepository struct {
	db  *gorm.DB
	key []byte
}

func NewUserRepository(db *gorm.DB, key []byte) UserRepository {
	return &userRepository{db: db, key: key}
}

func (r *userRepository) Create(user *dbmodels.User) error {
	if user.CreditUserPassword != nil {
		encryptedPassword, err := utils.Encrypt(*user.CreditUserPassword, r.key)
		if err != nil {
			return err
		}
		user.CreditUserPassword = &encryptedPassword
	}

	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*dbmodels.User, error) {
	var user dbmodels.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.CreditUserPassword != nil {
		decryptedPassword, err := utils.Decrypt(*user.CreditUserPassword, r.key)
		if err != nil {
			return nil, err
		}
		user.CreditUserPassword = &decryptedPassword
	}

	return &user, nil
}

func (r *userRepository) Update(user *dbmodels.User) error {
	if user.CreditUserPassword != nil {
		encryptedPassword, err := utils.Encrypt(*user.CreditUserPassword, r.key)
		if err != nil {
			return err
		}
		user.CreditUserPassword = &encryptedPassword
	}

	return r.db.Save(user).Error
}

func (r *userRepository) DeleteByEmail(email string) error {
	return r.db.Where("email = ?", email).Delete(&dbmodels.User{}).Error
}
