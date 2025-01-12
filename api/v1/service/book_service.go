package service

import (
	"fmt"
	"library_management_system/api/v1/models"
)

// BookService interface
type BookService interface {
	CreateBook(book *models.Books) (*models.Books, error)
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

	fmt.Println("bookDetail: ", bookDetail)
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
