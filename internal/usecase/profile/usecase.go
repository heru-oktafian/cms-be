package profile

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type Usecase struct {
	repo repository.ProfileRepository
}

func NewUsecase(repo repository.ProfileRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) Get() (*entity.Profile, error) {
	return u.repo.Get()
}

func (u *Usecase) Upsert(input *entity.Profile) (*entity.Profile, error) {
	return u.repo.Upsert(input)
}
