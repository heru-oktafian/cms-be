package repository

import "github.com/heru-oktafian/cms-be/internal/domain/entity"

type ProfileRepository interface {
	Get() (*entity.Profile, error)
	Upsert(profile *entity.Profile) (*entity.Profile, error)
}
