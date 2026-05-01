package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type ExperienceRepository interface {
	ListActive() ([]entity.Experience, error)
	List() ([]entity.Experience, error)
	GetByID(id uint) (*entity.Experience, error)
	Create(experience *entity.Experience) (*entity.Experience, error)
	Update(experience *entity.Experience) (*entity.Experience, error)
	Delete(id uint) error
}
