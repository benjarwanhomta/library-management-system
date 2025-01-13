package repo

import (
	"library_management_system/api/v1/models"
	"library_management_system/api/v1/service"
	"log"

	"gorm.io/gorm"
)

// BorrowingRecords represents BorrowingRecords in the library
type BorrowingRecords struct {
	RecordID   int    `json:"record_id"`
	BookID     int    `json:"book_id"`
	MemberID   int    `json:"member_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`
	Status     string `json:"status"`
}

// borrowingRecordsRepo implementation
type borrowingRecordsRepo struct {
	DB *gorm.DB
}

// NewBorrowingRecords creates a new BooksRepository
func NewBorrowingRecordsRepo(db *gorm.DB) service.BorrowsRepository {
	return &borrowingRecordsRepo{DB: db}
}

// GetTopBorrowingRecords implements service.BorrowsRepository.
func (b *borrowingRecordsRepo) GetTopBorrowingRecords() ([]models.TopBorrowedBook, error) {
	// Prepare the query to get the top borrowed books
	var topBooks []models.TopBorrowedBook

	result := b.DB.Raw(`
		SELECT b.title, COUNT(br.book_id) AS borrow_count
		FROM borrowingrecords br
		JOIN Books b ON br.book_id = b.book_id
		WHERE br.status = 'borrowed'
		GROUP BY br.book_id
		ORDER BY borrow_count DESC
		LIMIT 10;
	`).Scan(&topBooks)
	if result.Error != nil {
		log.Printf("GetTopBorrowingRecords has err: %s", result.Error.Error())
		return nil, result.Error
	}

	// If no data is found, return an empty slice
	if len(topBooks) == 0 {
		return []models.TopBorrowedBook{}, nil
	}

	// Return the result with the top borrowed books
	return topBooks, nil
}
