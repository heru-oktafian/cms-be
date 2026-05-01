package skill

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type Usecase struct {
	repo repository.SkillRepository
}

func NewUsecase(repo repository.SkillRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) ListActive() ([]entity.Skill, error) {
	return u.repo.ListActive()
}

func (u *Usecase) List() ([]entity.Skill, error) {
	return u.repo.List()
}

func (u *Usecase) GetByID(id uint) (*entity.Skill, error) {
	return u.repo.GetByID(id)
}

func (u *Usecase) Create(skill *entity.Skill) (*entity.Skill, error) {
	return u.repo.Create(skill)
}

func (u *Usecase) Update(skill *entity.Skill) (*entity.Skill, error) {
	return u.repo.Update(skill)
}

func (u *Usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
