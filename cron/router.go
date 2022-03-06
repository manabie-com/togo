package cron

import (
	cron "github.com/robfig/cron/v3"
)

// InitRouter initialize routing information
func InitRouter() {
	c := cron.New()
	c.Start()
}
