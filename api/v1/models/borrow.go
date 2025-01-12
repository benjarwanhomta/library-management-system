package models

// BorrowingRecords represents a record of a borrowing book
type BorrowingRecords struct {
	RecordID   int    `json:"record_id"`
	BookID     int    `json:"book_id"`
	MemberID   int    `json:"member_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`
}
