package contact_message

import (
	"strings"
	"time"

	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
)

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

type Usecase struct {
	repo repository.ContactMessageRepository
}

func NewUsecase(repo repository.ContactMessageRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) List() ([]entity.ContactMessage, error) {
	return u.repo.List()
}

func (u *Usecase) GetByID(id uint) (*entity.ContactMessage, error) {
	return u.repo.GetByID(id)
}

func (u *Usecase) Create(message *entity.ContactMessage) (*entity.ContactMessage, error) {
	message.Name = strings.TrimSpace(message.Name)
	message.Email = strings.TrimSpace(message.Email)
	message.Phone = strings.TrimSpace(message.Phone)
	message.Subject = strings.TrimSpace(message.Subject)
	message.Message = strings.TrimSpace(message.Message)
	message.Status = StatusNew
	message.FollowUpNotes = ""
	message.FollowedUpAt = nil
	return u.repo.Create(message)
}

func (u *Usecase) Update(message *entity.ContactMessage) (*entity.ContactMessage, error) {
	message.Name = strings.TrimSpace(message.Name)
	message.Email = strings.TrimSpace(message.Email)
	message.Phone = strings.TrimSpace(message.Phone)
	message.Subject = strings.TrimSpace(message.Subject)
	message.Message = strings.TrimSpace(message.Message)
	message.FollowUpNotes = strings.TrimSpace(message.FollowUpNotes)

	now := time.Now()
	if message.Status == StatusInProgress || message.Status == StatusDone {
		message.FollowedUpAt = &now
	} else {
		message.FollowedUpAt = nil
	}

	return u.repo.Update(message)
}

func IsValidStatus(status string) bool {
	switch status {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}

func (u *Usecase) Count() (int, error) {
	return u.repo.Count()
}
