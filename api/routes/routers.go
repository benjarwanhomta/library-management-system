package routes

import (
	v1 "library_management_system/api/v1"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitialRoute(app *fiber.App, db *gorm.DB) {
	v1.RouteV1(app, db)
}
