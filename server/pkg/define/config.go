package define

import (
	"github.com/spf13/viper"
	"time"
)

func RateLimitCreateTaskDuration() time.Duration {
	return viper.GetDuration("ratelimit.createtask.duration")
}

func RateLimitCreateTaskTimes() int {
	return viper.GetInt("ratelimit.createtask.times")
}
