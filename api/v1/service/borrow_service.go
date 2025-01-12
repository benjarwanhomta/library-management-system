package service

import "library_management_system/api/v1/models"

// ตัวแปรสำหรับเก็บข้อมูลการยืม
var borrows = []models.BorrowingRecords{
	// {ID: 1, BookID: 1, UserID: 101, BorrowDate: "2025-01-01", ReturnDate: "2025-01-15"},
}

// BorrowBook borrows a book for a user
func BorrowBook(borrow models.BorrowingRecords) models.BorrowingRecords {
	borrows = append(borrows, borrow)
	return borrow
}

// ReturnBook marks a book as returned
func ReturnBook(id int) bool {
	for i, borrow := range borrows {
		if borrow.RecordID == id {
			borrows = append(borrows[:i], borrows[i+1:]...)
			return true
		}
	}
	return false
}
