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
	public.Get("/skills", h.Placeholder("public skills endpoint"))
	public.Get("/experiences", h.Placeholder("public experiences endpoint"))
	public.Get("/social-links", h.Placeholder("public social links endpoint"))

	adminAuth := api.Group("/admin/auth")
	adminAuth.Post("/login", h.LoginAdmin)

	admin := api.Group("/admin", middleware.AdminJWT(container.Config.JWTSecret))
	admin.Get("/profile", h.GetAdminProfile)
	admin.Put("/profile", h.UpsertAdminProfile)
	admin.Get("/projects", h.ListAdminProjects)
	admin.Get("/projects/:id", h.GetAdminProject)
	admin.Post("/projects", h.CreateAdminProject)
	admin.Put("/projects/:id", h.UpdateAdminProject)
	admin.Delete("/projects/:id", h.DeleteAdminProject)
	admin.Get("/skills", h.Placeholder("admin skills endpoint"))
	admin.Get("/experiences", h.Placeholder("admin experiences endpoint"))
	admin.Get("/social-links", h.Placeholder("admin social links endpoint"))

	return app
}
