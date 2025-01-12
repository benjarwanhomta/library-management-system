package repo

import (
	"errors"
	"library_management_system/api/v1/models"
	"library_management_system/api/v1/service"
	"log"

	"gorm.io/gorm"
)

// Books represents books in the library
type Books struct {
	BookID       int    `json:"book_id"`
	Title        string `json:"title"`
	AuthorID     int    `json:"author_id"`
	CategoryID   int    `json:"category_id"`
	ISBN         int    `json:"isbn"`
	Description  string `json:"description"`
	AvailableQTY int    `json:"available_qty"`
}

// bookRepo implementation
type bookRepo struct {
	DB *gorm.DB
}

// NewBooksRepo creates a new BooksRepository
func NewBooksRepo(db *gorm.DB) service.BooksRepository {
	return &bookRepo{DB: db}
}

// CreateBooks implements service.BooksRepository.
func (r *bookRepo) CreateBooks(book *models.Books) error {
	result := r.DB.Create(book)
	if result.Error != nil {
		log.Printf("CreateBooks is has error: %s", result.Error.Error())
		return result.Error
	}
	return nil
}

// GetSingleByBiikID implements service.BooksRepository.
func (r *bookRepo) GetSingleByTitle(title string) (*models.Books, error) {
	var book models.Books
	result := r.DB.Where(`title = ?`, title).First(&book)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("GetSingleByBookName is has error: %s", result.Error.Error())
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &book, nil
}
