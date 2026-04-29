package app

import (
	"log"

	"github.com/heru-oktafian/cms-be/internal/config"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	"github.com/heru-oktafian/cms-be/internal/infrastructure/database"
	"gorm.io/gorm"
)

type App struct {
	Config config.Config
	DB     *gorm.DB
}

func Bootstrap() *App {
	cfg := config.Load()
	db, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if err := db.AutoMigrate(
		&entity.Profile{},
		&entity.Project{},
		&entity.Skill{},
		&entity.Experience{},
		&entity.SocialLink{},
		&entity.AdminUser{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return &App{Config: cfg, DB: db}
}
