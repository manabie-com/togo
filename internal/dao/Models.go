package dao

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Con struct {
	db  *gorm.DB
	Ctx *fiber.Ctx
}
