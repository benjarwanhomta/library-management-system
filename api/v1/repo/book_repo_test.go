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

	// สร้าง repo ด้วย DB ที่ตั้งค่าแล้ว
	bookRepo := repo.NewBooksRepo(db)

	// ข้อมูลหนังสือที่ใช้ทดสอบ
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

	// ตรวจสอบว่า book ถูกบันทึกลงฐานข้อมูล
	var book models.Books
	result := db.Where("title = ?", "Test Book").First(&book)
	assert.NoError(t, result.Error)
	assert.Equal(t, "Test Book", book.Title)
}

func TestGetSingleByTitle(t *testing.T) {

	// สร้าง repo ด้วย DB ที่ตั้งค่าแล้ว
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
