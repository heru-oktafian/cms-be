package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ContactMessageRepository struct {
	db *gorm.DB
}

func NewContactMessageRepository(db *gorm.DB) *ContactMessageRepository {
	return &ContactMessageRepository{db: db}
}

func (r *ContactMessageRepository) List() ([]entity.ContactMessage, error) {
	var messages []entity.ContactMessage
	if err := r.db.Order("created_at DESC, id DESC").Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *ContactMessageRepository) GetByID(id uint) (*entity.ContactMessage, error) {
	var message entity.ContactMessage
	if err := r.db.First(&message, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &message, nil
}

func (r *ContactMessageRepository) Create(message *entity.ContactMessage) (*entity.ContactMessage, error) {
	if err := r.db.Create(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *ContactMessageRepository) Update(message *entity.ContactMessage) (*entity.ContactMessage, error) {
	if err := r.db.Save(message).Error; err != nil {
		return nil, err
	}
	return message, nil
}
