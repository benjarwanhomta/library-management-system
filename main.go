package main

import (
	"library_management_system/api/routes"
	"library_management_system/config"
	_ "library_management_system/docs" // เชื่อมต่อกับไฟล์ docs ที่เราจะสร้าง
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// เชื่อมต่อกับฐานข้อมูล
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// ใช้งานฐานข้อมูล (เช่น migration หรือ query)
	log.Println("Database connection successful")

	routes.InitialRoute(app, db)

	// Start the server
	app.Listen(":3000")

}
