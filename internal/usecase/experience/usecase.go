package experience

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type Usecase struct {
	repo repository.ExperienceRepository
}

func NewUsecase(repo repository.ExperienceRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List() ([]entity.Experience, error) {
	return u.repo.List()
}

func (u *Usecase) GetByID(id uint) (*entity.Experience, error) {
	return u.repo.GetByID(id)
}

func (u *Usecase) Create(experience *entity.Experience) (*entity.Experience, error) {
	return u.repo.Create(experience)
}

func (u *Usecase) Update(experience *entity.Experience) (*entity.Experience, error) {
	return u.repo.Update(experience)
}

func (u *Usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
