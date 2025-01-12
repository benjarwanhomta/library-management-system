package service

import (
	"errors"
	"library_management_system/api/v1/models"

	"gorm.io/gorm"
)

// BookService interface
type BookService interface {
	CreateBook(book *models.Books) (*models.Books, error)
	UpdateBook(book *models.Books) (*models.Books, error)
	DeleteBookByBookID(bookID int) error
	GetDetailBookByBookID(bookID int) (*models.Books, error)
	GetSearchBookAll(bookID int) ([]models.Books, error)
	GetTopBorrowedBook(bookID int) ([]models.Books, error)
}

// bookService implementation
type bookService struct {
	bookRepo BooksRepository
}

// NewBooksService creates a new BookService
func NewBooksService(bookRepo BooksRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

// CreateBook implements BooksService.
func (s *bookService) CreateBook(book *models.Books) (*models.Books, error) {
	// get the book
	bookDetail, err := s.bookRepo.GetSingleByTitle(book.Title)
	if err != nil {
		return nil, err
	}

	if bookDetail == nil {
		err := s.bookRepo.CreateBooks(book)
		if err != nil {
			return nil, err
		}

		// get the book
		bookDetail, err = s.bookRepo.GetSingleByTitle(book.Title)
		if err != nil {
			return nil, err
		}

	}

	return bookDetail, nil
}

// DeleteBookByBookID implements BookService.
func (s *bookService) DeleteBookByBookID(bookID int) error {
	err := s.bookRepo.DeleteByBookID(bookID)
	if err != nil {
		return err
	}

	return nil
}

// EditBookByBookID implements BookService.
func (s *bookService) UpdateBook(book *models.Books) (*models.Books, error) {
	// get the book
	_, err := s.bookRepo.GetSingleByBookID(book.BookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetDetailNotFound
		}
		return nil, err
	}

	bookLastUpdate, err := s.bookRepo.UpdateBook(book)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetDetailNotFound
		}
		return nil, err
	}

	return bookLastUpdate, nil

}

// GetDetailBookByBookID implements BookService.
func (s *bookService) GetDetailBookByBookID(bookID int) (*models.Books, error) {
	// get the book
	bookDetail, err := s.bookRepo.GetSingleByBookID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGetDetailNotFound
		}
		return nil, err
	}

	return bookDetail, nil
}

// GetSearchBookAll implements BookService.
func (s *bookService) GetSearchBookAll(bookID int) ([]models.Books, error) {
	panic("unimplemented")
}

// GetTopBorrowedBook implements BookService.
func (s *bookService) GetTopBorrowedBook(bookID int) ([]models.Books, error) {
	panic("unimplemented")
}
