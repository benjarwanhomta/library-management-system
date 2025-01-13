package models

// Borrowingrecords represents a record of a borrowing book
type Borrowingrecords struct {
	RecordID   int    `json:"record_id"`
	BookID     int    `json:"book_id"`
	MemberID   int    `json:"member_id"`
	BorrowDate string `json:"borrow_date"`
	ReturnDate string `json:"return_date"`
	Status     string `json:"status"`
}

// TopBorrowedBook struct
type TopBorrowedBook struct {
	Title       string `json:"title"`
	BorrowCount int    `json:"borrow_count"`
}
