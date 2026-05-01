package tool

import (
	"strings"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

type UseCase struct {
	repo repository.ToolRepository
}

func NewUseCase(repo repository.ToolRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (u *UseCase) GetAllActive() ([]entity.Tool, error) {
	return u.repo.GetAllActive()
}

func (u *UseCase) GetAll() ([]entity.Tool, error) {
	return u.repo.GetAll()
}

func (u *UseCase) Create(input *entity.Tool) error {
	input.Name = strings.TrimSpace(input.Name)
	input.IconPath = strings.TrimSpace(input.IconPath)
	input.URL = strings.TrimSpace(input.URL)
	return u.repo.Create(input)
}

func (u *UseCase) Update(input *entity.Tool) error {
	input.Name = strings.TrimSpace(input.Name)
	input.IconPath = strings.TrimSpace(input.IconPath)
	input.URL = strings.TrimSpace(input.URL)
	return u.repo.Update(input)
}

func (u *UseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}
