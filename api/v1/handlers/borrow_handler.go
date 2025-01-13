package handlers

import (
	"library_management_system/api/v1/service"

	"github.com/gofiber/fiber/v2"
)

// BorrowHandler struct
type BorrowHandler struct {
	borrowService service.BorrowService
}

// NewBookHandler creates a new BookHandler
func NewBorrowHandler(borrowService service.BorrowService) *BorrowHandler {
	return &BorrowHandler{borrowService: borrowService}
}

// TopBorrowedBook handles Get request to get a top borrowed book.
func (h *BorrowHandler) TopBorrowedBook(c *fiber.Ctx) error {

	topBorrowed, err := h.borrowService.GetTopBorrowedBook()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get top borrowed book",
			"status":  "Fail",
			"value":   nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "Successfully get top borrowed book",
			"status":  "Success",
			"value":   topBorrowed,
		})
}
