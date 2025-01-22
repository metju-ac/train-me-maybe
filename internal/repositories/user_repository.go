package repositories

import (
	"fmt"

	"github.com/metju-ac/train-me-maybe/internal/dbmodels"
	"github.com/metju-ac/train-me-maybe/internal/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Key []byte
}

func (r *UserRepository) Create(user *dbmodels.User) error {
	if user.CreditUserPassword != nil {
		encryptedPassword, err := utils.Encrypt(*user.CreditUserPassword, r.Key)
		if err != nil {
			return fmt.Errorf("failed to encrypt password: %w", err)
		}
		user.CreditUserPassword = &encryptedPassword
	}

	return r.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*dbmodels.User, error) {
	var user dbmodels.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.CreditUserPassword != nil {
		decryptedPassword, err := utils.Decrypt(*user.CreditUserPassword, r.Key)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt password: %w", err)
		}
		user.CreditUserPassword = &decryptedPassword
	}

	return &user, nil
}

func (r *UserRepository) Update(user *dbmodels.User) error {
	if user.CreditUserPassword != nil {
		encryptedPassword, err := utils.Encrypt(*user.CreditUserPassword, r.Key)
		if err != nil {
			return fmt.Errorf("failed to encrypt password: %w", err)
		}
		user.CreditUserPassword = &encryptedPassword
	}

	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteByEmail(email string) error {
	return r.DB.Where("email = ?", email).Delete(&dbmodels.User{}).Error
}
