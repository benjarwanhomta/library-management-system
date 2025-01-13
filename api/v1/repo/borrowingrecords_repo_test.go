package repo_test

import (
	"fmt"
	"library_management_system/api/v1/models"
	"library_management_system/api/v1/repo"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConfig *gorm.DB

func init() {
	user := "root"
	password := "root"
	dbname := "library_management_system"
	host := "127.0.0.1"
	port := "3306"

	// สร้าง DSN (Data Source Name) สำหรับการเชื่อมต่อ
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// เชื่อมต่อฐานข้อมูล
	dbOpen, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database is err: ", err)
	}

	dbConfig = dbOpen
}

// Test the GetTopBorrowingRecords function
func TestGetTopBorrowingRecords(t *testing.T) {
	borrowingRecordsRepo := repo.NewBorrowingRecordsRepo(dbConfig)

	// สร้างข้อมูลตัวอย่าง
	author := models.Authors{Name: "John Doe"}
	category := models.Categories{CategoryName: "Programming"}

	// บันทึกข้อมูลตัวอย่างลงในฐานข้อมูล
	dbConfig.Create(&author)
	dbConfig.Create(&category)

	dbConfig.Where(`name = ?`, author.Name).Find(&author)
	dbConfig.Where(`category_name = ?`, category.CategoryName).Find(&category)

	book := models.Books{
		Title:        "Go Programming",
		AuthorID:     author.AuthorID,
		CategoryID:   category.CategoryID,
		ISBN:         "123456",
		Description:  "A test book for integration test",
		AvailableQTY: 0,
	}

	dbConfig.Create(&book)

	dbConfig.Where(`title = ?`, book.Title).Find(&book)

	member := models.Members{
		Name:        "John Doe",
		Email:       "JohnDoe@mail.com",
		PhoneNumber: "0123456789",
		Address:     "789/10 ถนนพระราม 2 แขวงบางมด เขตจอมทอง กรุงเทพมหานคร",
	}

	dbConfig.Create(&member)
	dbConfig.Where(`name = ?`, member.Name).Find(&member)

	// Create borrowing records for testing
	borrowingRecord := models.Borrowingrecords{
		BookID:     book.BookID,
		MemberID:   member.MemberID,
		BorrowDate: "2025-01-10",
		ReturnDate: "2025-01-12",
		Status:     "borrowed",
	}

	err := dbConfig.Create(&borrowingRecord).Error
	if err != nil {
		t.Fatalf("Failed to insert borrowing record: %v", err)
	}

	dbConfig.Where(`book_id = ? AND member_id = ? AND status = "borrowed"`, borrowingRecord.BookID, borrowingRecord.MemberID).Find(&borrowingRecord)

	// Call the method to get top borrowed books
	topBooks, err := borrowingRecordsRepo.GetTopBorrowingRecords()

	// Assert no error and check that the result is not empty
	assert.Nil(t, err, "Expected no error, but got an error")
	assert.NotEmpty(t, topBooks, "Expected top borrowed books, but got none")

	// Assert that the top book matches the inserted book
	assert.Equal(t, "Go Programming", topBooks[0].Title, "Expected title to be 'Go Programming'")
	assert.Equal(t, 1, topBooks[0].BorrowCount, "Expected borrow count to be 1")

	dbConfig.Where(`record_id = ?`, borrowingRecord.RecordID).Delete(&borrowingRecord)
	dbConfig.Where(`book_id = ?`, book.BookID).Delete(&book)
	dbConfig.Where(`author_id = ?`, author.AuthorID).Delete(&author)
	dbConfig.Where(`category_id = ?`, category.CategoryID).Delete(&category)
	dbConfig.Where(`member_id = ?`, member.MemberID).Delete(&member)
}

// func TestTop(t *testing.T) {
// 	borrowingRecordsRepo := repo.NewBorrowingRecordsRepo(dbConfig)
// 	// Call the method to get top borrowed books
// 	topBooks, err := borrowingRecordsRepo.GetTopBorrowingRecords()

// 	// Assert no error and check that the result is not empty
// 	assert.Nil(t, err, "Expected no error, but got an error")
// 	assert.NotEmpty(t, topBooks, "Expected top borrowed books, but got none")

// 	// Assert that the top book matches the inserted book
// 	assert.Equal(t, "Go Programming", topBooks[0].Title, "Expected title to be 'Go Programming'")
// 	assert.Equal(t, 1, topBooks[0].BorrowCount, "Expected borrow count to be 1")

// }
