package repo

import (
	"errors"
	"fmt"
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

// GetSingleByBookID implements service.BooksRepository.
func (r *bookRepo) GetSingleByBookID(bookID int) (*models.Books, error) {
	var book models.Books
	result := r.DB.Preload("Author").Preload("Category").First(&book, bookID)
	if result.Error != nil {
		log.Printf("GetSingleByBookName is has error: %s", result.Error.Error())
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &book, nil
}

// DeleteByBookID implements service.BooksRepository.
func (r *bookRepo) DeleteByBookID(bookID int) error {
	var book models.Books

	result := r.DB.Where("book_id = ?", bookID).Delete(&book)
	if result.Error != nil {
		log.Printf("DeleteByBookID is has error: %s", result.Error.Error())
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("Book with ID %d not found", bookID)
	}

	return nil
}

func (r *bookRepo) UpdateBook(book *models.Books) (*models.Books, error) {
	result := r.DB.Model(&models.Books{}).Where("book_id = ?", book.BookID).Updates(book)
	if result.Error != nil {
		log.Printf("UpdateBook has error: %s", result.Error.Error())
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Book with ID %d not found", book.BookID)
	}

	return book, nil
}
