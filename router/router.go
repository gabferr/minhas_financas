package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()

	initializeRoutes(router)
	//Run Server
	router.Run() // listen and serve on 0.0.0.0:8080

}
