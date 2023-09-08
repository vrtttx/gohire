package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	r := gin.Default()

	// initialize routes
	initializeRoutes(r)

	// run the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}