package workers

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/jinzhu/gorm"

	"github.com/manabie-com/togo/internal/storages"
)

// CronResetUserMaxTaskEveryDay for user can create new tasks
func CronResetUserMaxTaskEveryDay(DB *gorm.DB) {
	fmt.Println("Run cronjob!!!!!")

	var resetMaxTask = func ()  {
		result := DB.Model(&storages.User{}).Update("current_number_task", 0)
		if result.Error != nil {
			fmt.Println("Reset max task get error: ",result.Error)
		}
		fmt.Println("Reset max task success!")
	}
	// Reset the data every 1 day at 00:00
	gocron.Every(1).Day().At("00:00").Do(resetMaxTask)
	<- gocron.Start()
}
