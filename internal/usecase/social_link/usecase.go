package social_link

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type Usecase struct {
	repo repository.SocialLinkRepository
}

func NewUsecase(repo repository.SocialLinkRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List() ([]entity.SocialLink, error) {
	return u.repo.List()
}

func (u *Usecase) GetByID(id uint) (*entity.SocialLink, error) {
	return u.repo.GetByID(id)
}

func (u *Usecase) Create(socialLink *entity.SocialLink) (*entity.SocialLink, error) {
	return u.repo.Create(socialLink)
}

func (u *Usecase) Update(socialLink *entity.SocialLink) (*entity.SocialLink, error) {
	return u.repo.Update(socialLink)
}

func (u *Usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
