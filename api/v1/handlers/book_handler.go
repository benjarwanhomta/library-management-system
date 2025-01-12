package handlers

import (
	"library_management_system/api/v1/models"
	"library_management_system/api/v1/service"

	"github.com/gofiber/fiber/v2"
)

// BookHandler struct
type BookHandler struct {
	bookService service.BookService
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

// CreateBook handles POST request to create a new book
func (h *BookHandler) CreateBook(c *fiber.Ctx) error {
	var bookReq models.Books
	if err := c.BodyParser(&bookReq); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookResp, err := h.bookService.CreateBook(&bookReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create book")
	}

	return c.Status(fiber.StatusOK).JSON(bookResp)
}
