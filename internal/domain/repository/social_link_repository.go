package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type SocialLinkRepository interface {
	ListActive() ([]entity.SocialLink, error)
	List() ([]entity.SocialLink, error)
	GetByID(id uint) (*entity.SocialLink, error)
	Create(socialLink *entity.SocialLink) (*entity.SocialLink, error)
	Update(socialLink *entity.SocialLink) (*entity.SocialLink, error)
	Delete(id uint) error
}
