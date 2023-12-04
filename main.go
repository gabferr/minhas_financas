package main

import (
	"fmt"

	"github.com/gabferr/minhas_financas/handler/config"
	"github.com/gabferr/minhas_financas/router"
)

func main() {
	//Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	//initialize router
	router.Initialize()
}
