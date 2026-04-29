package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type AdminUserRepository interface {
	GetByEmail(email string) (*entity.AdminUser, error)
	Create(user *entity.AdminUser) (*entity.AdminUser, error)
}
