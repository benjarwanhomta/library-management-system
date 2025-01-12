package routes

import (
	"library_management_system/api/v1/handlers"
	"library_management_system/api/v1/repo"
	"library_management_system/api/v1/service"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

// RouteV1 defines the routes for the v1 API version.
func RouteV1(app *fiber.App, db *gorm.DB) {
	// Initialize repositories, services, and handlers
	bookRepo := repo.NewBooksRepo(db)
	bookService := service.NewBooksService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	// Swagger UI route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Create a group for v1 routes
	v1 := app.Group("/api/v1")

	// Define routes for the v1 group
	v1.Post("/add_book", bookHandler.CreateBook)
}
