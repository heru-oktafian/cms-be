package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type ToolRepository interface {
	GetAllActive() ([]entity.Tool, error)
	GetAll() ([]entity.Tool, error)
	GetByID(id uint) (*entity.Tool, error)
	Create(tool *entity.Tool) error
	Update(tool *entity.Tool) error
	Delete(id uint) error
}
