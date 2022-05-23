package initializer

import (
	"github.com/manabie-com/togo/handlers"
	"github.com/manabie-com/togo/models"
	"gorm.io/gorm"
)

var GlobalConfig GlobalParams

func InitApplication(envFileName string) (err error) {
	// Init all Global parameters
	GlobalConfig = newGlobalParams()
	if err = GlobalConfig.LoadAllConfig(envFileName); err != nil {
		return
	}

	models := InitModels(GlobalConfig.Components.Db)
	handlers := InitHandlers(models)
	handlers.SetupRouter(GlobalConfig.Components.GinEngine)

	return nil
}

func InitModels(db *gorm.DB) *models.Models {
	return &models.Models{
		User: models.NewUserModel(db),
		Task: models.NewTaskModel(db),
	}
}

func InitHandlers(models *models.Models) *handlers.Handlers {
	return &handlers.Handlers{
		Task: handlers.NewTaskHandler(models),
	}
}
