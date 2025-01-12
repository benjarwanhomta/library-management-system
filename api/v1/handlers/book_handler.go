package handlers

import (
	"library_management_system/api/v1/models"
	"library_management_system/api/v1/service"
	"strconv"

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
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Failed to body parser create book",
			"status":  "Fail",
			"value":   nil,
		})
	}

	bookResp, err := h.bookService.CreateBook(&bookReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create book",
			"status":  "Fail",
			"value":   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Successfully created",
			"status":  "Success",
			"value":   bookResp,
		})
}

// EditBook handles PUT request to edit a book
func (h *BookHandler) EditBook(c *fiber.Ctx) error {
	// var bookReq models.Books
	// if err := c.BodyParser(&bookReq); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	// }

	// bookResp, err := h.bookService.CreateBook(&bookReq)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).SendString("Failed to create book")
	// }

	return c.Status(fiber.StatusOK).JSON(nil)
}

// DeleteBookByBookID handles Delete request to delete a book by book ID.
func (h *BookHandler) DeleteBookByBookID(c *fiber.Ctx) error {
	bookIDParam := c.Params("book_id")

	if bookIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "book_id is required",
			"status":  "Fail",
		})
	}

	// Convert bookIDParam to int
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid format for book_id",
			"status":  "Fail",
		})
	}

	// DeleteBookByBookID
	errDeleteBook := h.bookService.DeleteBookByBookID(bookID)
	if errDeleteBook != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": errDeleteBook.Error(),
			"status":  "Fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully delete book by book id",
		"status":  "Success",
	})
}

// DetailBookByBookID handles Get request to get a book by book ID.
func (h *BookHandler) DetailBookByBookID(c *fiber.Ctx) error {
	bookIDParam := c.Params("book_id")

	if bookIDParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "book_id is required",
			"status":  "Fail",
		})
	}

	// Convert bookIDParam to int
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid format for book_id",
			"status":  "Fail",
		})
	}

	// GetDetailBookByBookID
	bookResp, err := h.bookService.GetDetailBookByBookID(bookID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Fail",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully get detail book by book id",
		"status":  "Success",
		"value":   bookResp,
	})
}

// SearchBookAll handles Get request to get a search book all.
func (h *BookHandler) SearchBookAll(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(nil)
}

// TopBorrowedBook handles Get request to get a top borrowed book.
func (h *BookHandler) TopBorrowedBook(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(nil)
}
