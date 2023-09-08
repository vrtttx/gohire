package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// initialize router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// run the server
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}