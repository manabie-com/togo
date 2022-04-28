package model

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Connect struct {
	Con *gorm.DB
	Ctx *fiber.Ctx
}
