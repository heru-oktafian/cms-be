package project

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type Usecase struct {
	repo repository.ProjectRepository
}

func NewUsecase(repo repository.ProjectRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) ListActive() ([]entity.Project, error) {
	return u.repo.ListActive()
}

func (u *Usecase) List() ([]entity.Project, error) {
	return u.repo.List()
}

func (u *Usecase) GetByID(id uint) (*entity.Project, error) {
	return u.repo.GetByID(id)
}

func (u *Usecase) Create(project *entity.Project) (*entity.Project, error) {
	return u.repo.Create(project)
}

func (u *Usecase) Update(project *entity.Project) (*entity.Project, error) {
	return u.repo.Update(project)
}

func (u *Usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
