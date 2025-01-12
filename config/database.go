package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	dsn := "root:root@tcp(127.0.0.1:3306)/library_management_system"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

// สร้างฟังก์ชันที่ใช้ในการเชื่อมต่อกับฐานข้อมูล
func ConnectDB() (*gorm.DB, error) {
	// โหลดค่าจากไฟล์ .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// อ่านค่าจาก environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// สร้าง DSN (Data Source Name) สำหรับการเชื่อมต่อ
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)

	// เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// คืนค่าการเชื่อมต่อฐานข้อมูล
	return db, nil
}
