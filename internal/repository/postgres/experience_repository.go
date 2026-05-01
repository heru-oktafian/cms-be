package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ExperienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{db: db}
}

func (r *ExperienceRepository) ListActive() ([]entity.Experience, error) {
	var experiences []entity.Experience
	if err := r.db.Where("is_active = ?", true).Order("sort_order ASC, id DESC").Find(&experiences).Error; err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *ExperienceRepository) List() ([]entity.Experience, error) {
	var experiences []entity.Experience
	if err := r.db.Order("sort_order ASC, id DESC").Find(&experiences).Error; err != nil {
		return nil, err
	}
	return experiences, nil
}

func (r *ExperienceRepository) GetByID(id uint) (*entity.Experience, error) {
	var experience entity.Experience
	if err := r.db.First(&experience, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &experience, nil
}

func (r *ExperienceRepository) Create(experience *entity.Experience) (*entity.Experience, error) {
	if err := r.db.Create(experience).Error; err != nil {
		return nil, err
	}
	return experience, nil
}

func (r *ExperienceRepository) Update(experience *entity.Experience) (*entity.Experience, error) {
	if err := r.db.Save(experience).Error; err != nil {
		return nil, err
	}
	return experience, nil
}

func (r *ExperienceRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Experience{}, id).Error
}
