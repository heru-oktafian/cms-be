package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type SkillRepository interface {
	List() ([]entity.Skill, error)
	GetByID(id uint) (*entity.Skill, error)
	Create(skill *entity.Skill) (*entity.Skill, error)
	Update(skill *entity.Skill) (*entity.Skill, error)
	Delete(id uint) error
}
