package define

import (
	"github.com/spf13/viper"
	"time"
)

func RateLimitCreateTaskDuration() time.Duration {
	return viper.GetDuration("ratelimit.createtask_duration") * time.Hour * 24
}

func RateLimitCreateTaskTimes() int {
	return viper.GetInt("ratelimit.createtask_times")
}
