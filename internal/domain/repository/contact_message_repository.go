package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type ContactMessageRepository interface {
	List() ([]entity.ContactMessage, error)
	GetByID(id uint) (*entity.ContactMessage, error)
	Create(message *entity.ContactMessage) (*entity.ContactMessage, error)
	Update(message *entity.ContactMessage) (*entity.ContactMessage, error)
}
