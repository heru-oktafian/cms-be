package postgres

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ToolRepository struct {
	db *gorm.DB
}

func NewToolRepository(db *gorm.DB) *ToolRepository {
	return &ToolRepository{db: db}
}

func (r *ToolRepository) GetAllActive() ([]entity.Tool, error) {
	var tools []entity.Tool
	err := r.db.Where("is_active = ?", true).Order("sort_order asc, id asc").Find(&tools).Error
	return tools, err
}

func (r *ToolRepository) GetAll() ([]entity.Tool, error) {
	var tools []entity.Tool
	err := r.db.Order("sort_order asc, id asc").Find(&tools).Error
	return tools, err
}

func (r *ToolRepository) GetByID(id uint) (*entity.Tool, error) {
	var tool entity.Tool
	if err := r.db.First(&tool, id).Error; err != nil {
		return nil, err
	}
	return &tool, nil
}

func (r *ToolRepository) Create(tool *entity.Tool) error {
	return r.db.Create(tool).Error
}

func (r *ToolRepository) Update(tool *entity.Tool) error {
	return r.db.Save(tool).Error
}

func (r *ToolRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Tool{}, id).Error
}
