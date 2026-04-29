package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type AdminUserRepository struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) *AdminUserRepository {
	return &AdminUserRepository{db: db}
}

func (r *AdminUserRepository) GetByEmail(email string) (*entity.AdminUser, error) {
	var user entity.AdminUser
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *AdminUserRepository) Create(user *entity.AdminUser) (*entity.AdminUser, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
