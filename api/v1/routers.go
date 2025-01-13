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

	borrowingRecordsRepo := repo.NewBorrowingRecordsRepo(db)
	borrowService := service.NewBorrowService(borrowingRecordsRepo)
	borrowHandler := handlers.NewBorrowHandler(borrowService)

	// Swagger UI route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Create a group for v1 routes
	v1 := app.Group("/api/v1")

	// Define routes for the v1 group
	v1.Post("/add_book", bookHandler.CreateBook)
	v1.Put("/edit_book", bookHandler.EditBook)
	v1.Delete("/delete_book/:book_id", bookHandler.DeleteBookByBookID)
	v1.Get("/detail_book/:book_id", bookHandler.DetailBookByBookID)
	v1.Get("/search_book_all", bookHandler.SearchBookAll)
	v1.Get("/top_borrowed_book", borrowHandler.TopBorrowedBook)
}
