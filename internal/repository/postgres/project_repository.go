package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) ListActive() ([]entity.Project, error) {
	var projects []entity.Project
	if err := r.db.Where("is_active = ?", true).Order("sort_order ASC, id DESC").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) List() ([]entity.Project, error) {
	var projects []entity.Project
	if err := r.db.Order("sort_order ASC, id DESC").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) GetByID(id uint) (*entity.Project, error) {
	var project entity.Project
	if err := r.db.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) Create(project *entity.Project) (*entity.Project, error) {
	if err := r.db.Create(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *ProjectRepository) Update(project *entity.Project) (*entity.Project, error) {
	if err := r.db.Save(project).Error; err != nil {
		return nil, err
	}
	return project, nil
}

func (r *ProjectRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Project{}, id).Error
}
