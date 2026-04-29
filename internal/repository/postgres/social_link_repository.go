package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type SocialLinkRepository struct {
	db *gorm.DB
}

func NewSocialLinkRepository(db *gorm.DB) *SocialLinkRepository {
	return &SocialLinkRepository{db: db}
}

func (r *SocialLinkRepository) List() ([]entity.SocialLink, error) {
	var socialLinks []entity.SocialLink
	if err := r.db.Order("sort_order ASC, id DESC").Find(&socialLinks).Error; err != nil {
		return nil, err
	}
	return socialLinks, nil
}

func (r *SocialLinkRepository) GetByID(id uint) (*entity.SocialLink, error) {
	var socialLink entity.SocialLink
	if err := r.db.First(&socialLink, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &socialLink, nil
}

func (r *SocialLinkRepository) Create(socialLink *entity.SocialLink) (*entity.SocialLink, error) {
	if err := r.db.Create(socialLink).Error; err != nil {
		return nil, err
	}
	return socialLink, nil
}

func (r *SocialLinkRepository) Update(socialLink *entity.SocialLink) (*entity.SocialLink, error) {
	if err := r.db.Save(socialLink).Error; err != nil {
		return nil, err
	}
	return socialLink, nil
}

func (r *SocialLinkRepository) Delete(id uint) error {
	return r.db.Delete(&entity.SocialLink{}, id).Error
}
