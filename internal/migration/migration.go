package migration

import (
	"fmt"
	"time"

	"github.com/TrinhTrungDung/togo/config"
	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/db"
	"github.com/TrinhTrungDung/togo/pkg/migration"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type Base struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Run() (resErr error) {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := db.New(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSslMode), false)
	if err != nil {
		panic(err)
	}

	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				resErr = fmt.Errorf("%s", x)
			case error:
				resErr = x
			default:
				resErr = fmt.Errorf("Unknown error: %+v", x)
			}
		}
	}()

	// Create migrations table to store migration version history
	initSQL := "CREATE TABLE IF NOT EXISTS migrations (id VARCHAR(255) PRIMARY KEY);"
	if err := db.Exec(initSQL).Error; err != nil {
		return err
	}

	migration.Run(db, []*gormigrate.Migration{
		{
			ID: "202204021000",
			Migrate: func(tx *gorm.DB) error {
				type Plan struct {
					Base
					Name     string `gorm:"type:varchar(30)"`
					MaxTasks int
				}

				type User struct {
					Base
					FirstName string `gorm:"type:varchar(255)"`
					LastName  string `gorm:"type:varchar(255)"`
					Email     string `gorm:"type:varchar(254);unique_index;not null"`
					Username  string `gorm:"type:varchar(255);unique_index;not null"`
					Password  string `gorm:"type:varchar(255);not null"`
				}

				type Task struct {
					Base
					Content string `gorm:"type:text"`
					UserID  int
				}

				type Subscription struct {
					UserID  int `gorm:"primary_key;autoIncrement:false"`
					PlanID  int `gorm:"primary_key;autoIncrement:false"`
					StartAt time.Time
					EndAt   time.Time
				}

				if err := tx.AutoMigrate(&User{}, &Plan{}, &Task{}, &Subscription{}); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("users", "plans", "tasks", "subscriptions")
			},
		},
		{
			ID: "202204031021",
			Migrate: func(tx *gorm.DB) error {
				defaultPlans := []*model.Plan{
					{
						Name:     model.FreemiumPlan,
						MaxTasks: 1,
					},
					{
						Name:     model.SilverPlan,
						MaxTasks: 10,
					},
					{
						Name:     model.GoldPlan,
						MaxTasks: 100,
					},
				}

				for _, plan := range defaultPlans {
					if err := tx.Create(plan).Error; err != nil {
						return err
					}
				}

				return nil
			},
		},
		{
			ID: "202204031056",
			Migrate: func(tx *gorm.DB) error {
				type Subscription struct {
					UserID  int `gorm:"primary_key;autoIncrement:false"`
					PlanID  int `gorm:"primary_key;autoIncrement:false"`
					StartAt time.Time
					EndAt   *time.Time
				}

				// Drop current subscriptions table
				// NOTE: Only do this when there's no existing subscription in database yet
				if err := tx.Migrator().DropTable("subscriptions"); err != nil {
					return err
				}

				if err := tx.AutoMigrate(&Subscription{}); err != nil {
					return err
				}

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("subscriptions")
			},
		},
	})

	return nil
}
