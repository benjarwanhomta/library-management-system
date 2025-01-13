package service

import (
	"errors"
	"library_management_system/api/v1/models"

	"gorm.io/gorm"
)

// BorrowService interface
type BorrowService interface {
	GetTopBorrowedBook() ([]models.TopBorrowedBook, error)
}

// borrowService implementation
type borrowService struct {
	borrowsRepo BorrowsRepository
}

// NewBorrowService creates a new BorrowService
func NewBorrowService(borrowsRepo BorrowsRepository) BorrowService {
	return &borrowService{borrowsRepo: borrowsRepo}
}

// GetTopBorrowedBook implements BorrowService.
func (s *borrowService) GetTopBorrowedBook() ([]models.TopBorrowedBook, error) {
	// get top borrow the book
	topBorrow, err := s.borrowsRepo.GetTopBorrowingRecords()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetDetailNotFound
		}
		return nil, err
	}

	return topBorrow, nil
}
