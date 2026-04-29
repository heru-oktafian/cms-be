package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	appctx "github.com/heru-oktafian/cms-be/internal/app"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	profileRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	profileUsecase "github.com/heru-oktafian/cms-be/internal/usecase/profile"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type Handler struct {
	app            *appctx.App
	profileUsecase *profileUsecase.Usecase
}

func NewHandler(app *appctx.App) *Handler {
	profileRepository := profileRepo.NewProfileRepository(app.DB)
	return &Handler{
		app:            app,
		profileUsecase: profileUsecase.NewUsecase(profileRepository),
	}
}

func (h *Handler) GetPublicProfile(c *fiber.Ctx) error {
	profile, err := h.profileUsecase.Get()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch profile", nil)
	}
	if profile == nil {
		return response.JSON(c, http.StatusOK, "profile not set", fiber.Map{})
	}
	return response.JSON(c, http.StatusOK, "ok", profile)
}

func (h *Handler) GetAdminProfile(c *fiber.Ctx) error {
	return h.GetPublicProfile(c)
}

func (h *Handler) UpsertAdminProfile(c *fiber.Ctx) error {
	var payload entity.Profile
	if err := c.BodyParser(&payload); err != nil {
		return response.JSON(c, http.StatusBadRequest, "invalid payload", nil)
	}
	profile, err := h.profileUsecase.Upsert(&payload)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to save profile", nil)
	}
	return response.JSON(c, http.StatusOK, "profile saved", profile)
}
