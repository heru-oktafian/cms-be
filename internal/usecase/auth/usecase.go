package auth

import (
	"errors"
	"strings"

	jwtAuth "github.com/heru-oktafian/cms-be/internal/infrastructure/auth"
	"github.com/heru-oktafian/cms-be/internal/config"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type Usecase struct {
	repo repository.AdminUserRepository
	cfg  config.Config
}

type LoginResult struct {
	Token string            `json:"token"`
	User  *entity.AdminUser `json:"user"`
}

func NewUsecase(repo repository.AdminUserRepository, cfg config.Config) *Usecase {
	return &Usecase{repo: repo, cfg: cfg}
}

func (u *Usecase) SeedDefaultAdmin(name, email, password string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || strings.TrimSpace(password) == "" {
		return nil
	}
	existing, err := u.repo.GetByEmail(email)
	if err != nil {
		return err
	}
	if existing != nil {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = u.repo.Create(&entity.AdminUser{
		Name:         strings.TrimSpace(name),
		Email:        email,
		PasswordHash: string(hash),
		IsActive:     true,
	})
	return err
}

func (u *Usecase) Login(email, password string) (*LoginResult, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil || !user.IsActive {
		return nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	signedToken, err := jwtAuth.GenerateAdminJWT(user, u.cfg.JWTSecret)
	if err != nil {
		return nil, err
	}
	return &LoginResult{Token: signedToken, User: user}, nil
}
