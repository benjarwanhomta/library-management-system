package service

import "library_management_system/api/v1/models"

//mockgen -source=D:/GoWorkspace/src/library_management_system/api/v1/service/borrow.go -destination=mocks/borrow_mock.go -package=mocks

// BorrowsRepository  interface
type BorrowsRepository interface {
	GetTopBorrowingRecords() ([]models.TopBorrowedBook, error)
}
