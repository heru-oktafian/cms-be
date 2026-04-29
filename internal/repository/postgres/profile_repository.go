package postgres

import (
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Get() (*entity.Profile, error) {
	var profile entity.Profile
	if err := r.db.First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *ProfileRepository) Upsert(profile *entity.Profile) (*entity.Profile, error) {
	current, err := r.Get()
	if err != nil {
		return nil, err
	}
	if current != nil {
		profile.ID = current.ID
	}
	if err := r.db.Save(profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}
