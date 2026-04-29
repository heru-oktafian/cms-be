package postgres

import (
	"errors"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type SkillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) List() ([]entity.Skill, error) {
	var skills []entity.Skill
	if err := r.db.Order("sort_order ASC, id DESC").Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *SkillRepository) GetByID(id uint) (*entity.Skill, error) {
	var skill entity.Skill
	if err := r.db.First(&skill, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &skill, nil
}

func (r *SkillRepository) Create(skill *entity.Skill) (*entity.Skill, error) {
	if err := r.db.Create(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *SkillRepository) Update(skill *entity.Skill) (*entity.Skill, error) {
	if err := r.db.Save(skill).Error; err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *SkillRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Skill{}, id).Error
}
