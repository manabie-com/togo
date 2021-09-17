package routers

import (
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"gorm.io/gorm"
)

// InitRouter initialize routing information
func InitRouter(db *gorm.DB) {
	// Initialize repository
	repo := repositories.InitRepositoryFactory(db)
	// Initialize service
	service := services.InitServiceFactory(repo)
}
