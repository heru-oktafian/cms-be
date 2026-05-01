package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type ProjectRepository interface {
	ListActive() ([]entity.Project, error)
	List() ([]entity.Project, error)
	GetByID(id uint) (*entity.Project, error)
	Create(project *entity.Project) (*entity.Project, error)
	Update(project *entity.Project) (*entity.Project, error)
	Delete(id uint) error
}
