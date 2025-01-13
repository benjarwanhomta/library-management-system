package repo_test

import (
	"fmt"
	"log"
	"testing"

	"library_management_system/api/v1/models"
	"library_management_system/api/v1/repo"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

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

	db = dbOpen
}

func TestCreateBooks(t *testing.T) {

	bookRepo := repo.NewBooksRepo(db)

	// ข้อมูลหนังสือที่ใช้ทดสอบ
	bookReq := &models.Books{
		Title:        "Test Book",
		AuthorID:     1,
		CategoryID:   1,
		ISBN:         "1234567890",
		Description:  "A test book for integration test",
		AvailableQTY: 10,
	}

	// ทดสอบการสร้างหนังสือ
	err := bookRepo.CreateBooks(bookReq)
	assert.NoError(t, err)

	// ตรวจสอบว่า book ถูกบันทึกลงฐานข้อมูล
	var book models.Books
	result := db.Where("title = ?", "Test Book").First(&book)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Test Book", book.Title)
}

func TestGetSingleByTitle(t *testing.T) {

	bookRepo := repo.NewBooksRepo(db)

	// สร้างข้อมูลหนังสือ
	bookReq := &models.Books{
		Title:        "Test Book",
		AuthorID:     1,
		CategoryID:   1,
		ISBN:         "123456789",
		Description:  "A test book for integration test",
		AvailableQTY: 10,
	}

	// ทดสอบการสร้างหนังสือ
	err := bookRepo.CreateBooks(bookReq)
	assert.NoError(t, err)

	// ทดสอบการดึงข้อมูลหนังสือตามชื่อ
	book, err := bookRepo.GetSingleByTitle("Test Book")
	assert.NoError(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, "Test Book", book.Title)

	// ทดสอบกรณีที่ไม่พบข้อมูล
	bookNotFound, err := bookRepo.GetSingleByTitle("Nonexistent Book")
	assert.NoError(t, err)
	assert.Nil(t, bookNotFound)
}

func TestGetSingleByBookID(t *testing.T) {

	bookRepo := repo.NewBooksRepo(db)

	// have data
	t.Run("Book exists", func(t *testing.T) {
		result, err := bookRepo.GetSingleByBookID(2)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 2, result.BookID)
		assert.Equal(t, "หนังสือ A", result.Title)
	})

	// not found
	t.Run("Book not found", func(t *testing.T) {
		result, err := bookRepo.GetSingleByBookID(0)
		assert.NoError(t, err)
		assert.Nil(t, result)
	})
}

func TestDeleteByBookID(t *testing.T) {

	bookRepo := repo.NewBooksRepo(db)

	// have data
	t.Run("Book exists", func(t *testing.T) {
		err := bookRepo.DeleteByBookID(4)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	// not found
	t.Run("Book not found", func(t *testing.T) {
		err := bookRepo.DeleteByBookID(0)
		assert.Error(t, err)
		assert.NotNil(t, err)
	})
}

func TestCreatdAndUpdateAndDeleteBook(t *testing.T) {
	bookRepo := repo.NewBooksRepo(db)

	// create a new book
	book := &models.Books{
		Title:        "Original Title",
		AuthorID:     1,
		CategoryID:   1,
		PublishYear:  2020,
		ISBN:         "9783-16-148450-0",
		Description:  "Test Book Description",
		AvailableQTY: 10,
	}

	err := bookRepo.CreateBooks(book)
	assert.NoError(t, err)

	result, err := bookRepo.GetSingleByTitle(book.Title)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// update
	updatedBook := &models.Books{
		BookID:       result.BookID, // ใช้ BookID ที่ได้จากการสร้างหนังสือ
		Title:        "Updated Title",
		AuthorID:     1,
		CategoryID:   1,
		PublishYear:  2021,
		ISBN:         "9783-16-148450-1",
		Description:  "Updated Book Description",
		AvailableQTY: 15,
	}

	// ทดสอบการอัปเดตหนังสือ
	updatedBookResult, err := bookRepo.UpdateBook(updatedBook)
	if err != nil {
		t.Fatalf("UpdateBook failed: %v", err)
	}

	assert.Equal(t, "Updated Title", updatedBookResult.Title)
	assert.Equal(t, 2021, updatedBookResult.PublishYear)
	assert.Equal(t, 15, updatedBookResult.AvailableQTY)

	resultBook, err := bookRepo.GetSingleByTitle(updatedBookResult.Title)
	assert.NoError(t, err)
	assert.NotNil(t, resultBook)

	assert.Equal(t, "Updated Title", resultBook.Title)
	assert.Equal(t, 2021, resultBook.PublishYear)
	assert.Equal(t, 15, resultBook.AvailableQTY)

	//delete
	err = bookRepo.DeleteByBookID(resultBook.BookID)
	assert.NoError(t, err)
	assert.Nil(t, err)
}

func TestGetAllByTitleAuthorCategory(t *testing.T) {
	bookRepo := repo.NewBooksRepo(db)

	// สร้างข้อมูลตัวอย่าง
	author := models.Authors{Name: "John Doe"}
	category := models.Categories{CategoryName: "Programming"}

	// บันทึกข้อมูลตัวอย่างลงในฐานข้อมูล
	db.Create(&author)
	db.Create(&category)

	db.Where(`name = ?`, author.Name).Find(&author)
	db.Where(`category_name = ?`, category.CategoryName).Find(&category)

	book := models.Books{
		Title:        "Go Programming",
		AuthorID:     author.AuthorID,
		CategoryID:   category.CategoryID,
		ISBN:         "123456",
		Description:  "A test book for integration test",
		AvailableQTY: 0,
	}

	db.Create(&book)

	db.Where(`title = ?`, book.Title).Find(&book)

	// ทดสอบฟังก์ชัน GetSingleByTitleAuthorCategory
	t.Run("Test with title", func(t *testing.T) {
		books, err := bookRepo.GetAllByTitleAuthorCategory("Go", "", "")
		assert.NoError(t, err)
		assert.Len(t, books, 1)
		assert.Equal(t, "Go Programming", books[0].Title)
	})

	t.Run("Test with author", func(t *testing.T) {
		books, err := bookRepo.GetAllByTitleAuthorCategory("", "John Doe", "")
		assert.NoError(t, err)
		assert.Len(t, books, 1)
		assert.Equal(t, author.AuthorID, books[0].AuthorID)
	})

	t.Run("Test with category", func(t *testing.T) {
		books, err := bookRepo.GetAllByTitleAuthorCategory("", "", "Programming")
		assert.NoError(t, err)
		assert.Len(t, books, 1)
		assert.Equal(t, category.CategoryID, books[0].CategoryID)
	})

	t.Run("Test with no results", func(t *testing.T) {
		books, err := bookRepo.GetAllByTitleAuthorCategory("Non-Existing Title", "Non-Existing Author", "Non-Existing Category")
		assert.NoError(t, err)
		assert.Len(t, books, 0)
	})

	db.Where(`book_id = ?`, book.BookID).Delete(&book)
	db.Where(`author_id = ?`, author.AuthorID).Delete(&author)
	db.Where(`category_id = ?`, category.CategoryID).Delete(&category)
}
