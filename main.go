package main

import (
	"github.com/gabferr/minhas_financas/config"
	"github.com/gabferr/minhas_financas/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//initialize router
	router.Initialize()
}
