package main

import (
	"github.com/vrtttx/gohire/config"
	"github.com/vrtttx/gohire/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("CONFIG_INIT_ERROR: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}