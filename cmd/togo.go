package main

import (
	"main/config"
	"main/internal/logger"
)

func main() {
	_ = config.Load()
	_ = logger.New()
}
