package service

import "library_management_system/api/v1/models"

//mockgen -source=D:/GoWorkspace/src/library_management_system/api/v1/service/book.go -destination=mocks/book_mock.go -package=mocks

// BooksRepository interface
type BooksRepository interface {
	CreateBooks(book *models.Books) error
	GetSingleByTitle(title string) (*models.Books, error)
	GetSingleByBookID(bookID int) (*models.Books, error)
	DeleteByBookID(bookID int) error
	UpdateBook(book *models.Books) (*models.Books, error)
}
