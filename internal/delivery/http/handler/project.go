package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/cms-be/internal/domain/entity"
	projectRepo "github.com/heru-oktafian/cms-be/internal/repository/postgres"
	projectUsecase "github.com/heru-oktafian/cms-be/internal/usecase/project"
	"github.com/heru-oktafian/cms-be/pkg/response"
)

type projectRequest struct {
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
	ThumbnailPath string `json:"thumbnail_path"`
	ProjectURL    string `json:"project_url"`
	RepoURL       string `json:"repo_url"`
	IsFeatured    bool   `json:"is_featured"`
	SortOrder     int    `json:"sort_order"`
}

func (h *Handler) initProjectUsecase() {
	if h.projectUsecase == nil {
		projectRepository := projectRepo.NewProjectRepository(h.app.DB)
		h.projectUsecase = projectUsecase.NewUsecase(projectRepository)
	}
}

func (h *Handler) ListPublicProjects(c *fiber.Ctx) error {
	h.initProjectUsecase()
	projects, err := h.projectUsecase.List()
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch projects", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", projects)
}

func (h *Handler) ListAdminProjects(c *fiber.Ctx) error {
	return h.ListPublicProjects(c)
}

func (h *Handler) GetAdminProject(c *fiber.Ctx) error {
	h.initProjectUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid project id", nil)
	}
	project, err := h.projectUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch project", nil)
	}
	if project == nil {
		return response.JSON(c, http.StatusNotFound, "project not found", nil)
	}
	return response.JSON(c, http.StatusOK, "ok", project)
}

func (h *Handler) CreateAdminProject(c *fiber.Ctx) error {
	h.initProjectUsecase()
	project, err := h.parseProjectRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	created, err := h.projectUsecase.Create(project)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to create project: %v", err), nil)
	}
	return response.JSON(c, http.StatusCreated, "project created", created)
}

func (h *Handler) UpdateAdminProject(c *fiber.Ctx) error {
	h.initProjectUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid project id", nil)
	}
	existing, err := h.projectUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch project", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "project not found", nil)
	}
	payload, err := h.parseProjectRequest(c)
	if err != nil {
		return response.JSON(c, http.StatusBadRequest, err.Error(), nil)
	}
	payload.ID = existing.ID
	payload.CreatedAt = existing.CreatedAt
	updated, err := h.projectUsecase.Update(payload)
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, fmt.Sprintf("failed to update project: %v", err), nil)
	}
	return response.JSON(c, http.StatusOK, "project updated", updated)
}

func (h *Handler) DeleteAdminProject(c *fiber.Ctx) error {
	h.initProjectUsecase()
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return response.JSON(c, http.StatusBadRequest, "invalid project id", nil)
	}
	existing, err := h.projectUsecase.GetByID(uint(id))
	if err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to fetch project", nil)
	}
	if existing == nil {
		return response.JSON(c, http.StatusNotFound, "project not found", nil)
	}
	if err := h.projectUsecase.Delete(uint(id)); err != nil {
		return response.JSON(c, http.StatusInternalServerError, "failed to delete project", nil)
	}
	return response.JSON(c, http.StatusOK, "project deleted", nil)
}

func (h *Handler) parseProjectRequest(c *fiber.Ctx) (*entity.Project, error) {
	var payload projectRequest
	if err := c.BodyParser(&payload); err != nil {
		return nil, err
	}

	payload.Title = strings.TrimSpace(payload.Title)
	payload.Slug = strings.TrimSpace(payload.Slug)
	payload.Summary = strings.TrimSpace(payload.Summary)
	payload.Description = strings.TrimSpace(payload.Description)
	payload.ThumbnailPath = strings.TrimSpace(payload.ThumbnailPath)
	payload.ProjectURL = strings.TrimSpace(payload.ProjectURL)
	payload.RepoURL = strings.TrimSpace(payload.RepoURL)

	if payload.Title == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "title is required")
	}
	if payload.Slug == "" {
		return nil, fiber.NewError(http.StatusBadRequest, "slug is required")
	}

	return &entity.Project{
		Title:         payload.Title,
		Slug:          payload.Slug,
		Summary:       payload.Summary,
		Description:   payload.Description,
		ThumbnailPath: payload.ThumbnailPath,
		ProjectURL:    payload.ProjectURL,
		RepoURL:       payload.RepoURL,
		IsFeatured:    payload.IsFeatured,
		SortOrder:     payload.SortOrder,
	}, nil
}
