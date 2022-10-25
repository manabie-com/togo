package console

import (
	"log"

	"github.com/robfig/cron"
)

func Schedule() {
	Job := cron.New()
	Job.AddFunc("30 * * * *", autoLoad)
	Job.Start()

	log.Printf("Scheduled jobs loaded")
}

func autoLoad() {
	log.Println("Schedule is running")
}
