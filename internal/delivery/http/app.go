package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/heru-oktafian/cms-be/internal/app"
	"github.com/heru-oktafian/cms-be/internal/delivery/http/handler"
	"github.com/heru-oktafian/cms-be/internal/delivery/http/middleware"
)

func NewApp(container *app.App) *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Static("/uploads", container.Config.UploadDir)

	h := handler.NewHandler(container)

	api := app.Group("/api/v1")
	api.Get("/health", h.Health)

	public := api.Group("/public")
	public.Get("/profile", h.GetPublicProfile)
	public.Get("/projects", h.ListPublicProjects)
	public.Get("/skills", h.ListPublicSkills)
	public.Get("/experiences", h.ListPublicExperiences)
	public.Get("/social-links", h.ListPublicSocialLinks)

	adminAuth := api.Group("/admin/auth")
	adminAuth.Post("/login", h.LoginAdmin)
	adminAuth.Post("/logout", h.LogoutAdmin)
	adminAuth.Get("/me", middleware.AdminJWT(container.Config.JWTSecret), h.GetAuthMe)

	admin := api.Group("/admin", middleware.AdminJWT(container.Config.JWTSecret))
	admin.Get("/profile", h.GetAdminProfile)
	admin.Put("/profile", h.UpsertAdminProfile)
	admin.Get("/projects", h.ListAdminProjects)
	admin.Get("/projects/:id", h.GetAdminProject)
	admin.Post("/projects", h.CreateAdminProject)
	admin.Put("/projects/:id", h.UpdateAdminProject)
	admin.Delete("/projects/:id", h.DeleteAdminProject)
	admin.Get("/skills", h.ListAdminSkills)
	admin.Get("/skills/:id", h.GetAdminSkill)
	admin.Post("/skills", h.CreateAdminSkill)
	admin.Put("/skills/:id", h.UpdateAdminSkill)
	admin.Delete("/skills/:id", h.DeleteAdminSkill)
	admin.Get("/experiences", h.ListAdminExperiences)
	admin.Get("/experiences/:id", h.GetAdminExperience)
	admin.Post("/experiences", h.CreateAdminExperience)
	admin.Put("/experiences/:id", h.UpdateAdminExperience)
	admin.Delete("/experiences/:id", h.DeleteAdminExperience)
	admin.Get("/social-links", h.ListAdminSocialLinks)
	admin.Get("/social-links/:id", h.GetAdminSocialLink)
	admin.Post("/social-links", h.CreateAdminSocialLink)
	admin.Put("/social-links/:id", h.UpdateAdminSocialLink)
	admin.Delete("/social-links/:id", h.DeleteAdminSocialLink)

	return app
}
