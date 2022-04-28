package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Con struct {
	Db  *gorm.DB
	Ctx *fiber.Ctx
}
