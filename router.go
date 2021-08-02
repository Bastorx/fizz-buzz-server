package main

import (
	"github.com/Bastorx/fizz-buzz-server/controllers"
	"github.com/gin-gonic/gin"
)

func initRouter(r* gin.Engine) {
	// PING
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// MAIN ROUTES
	r.POST("/fizz-buzz", controllers.FizzBuzz)

	// 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "NOT_FOUND", "message": "Page not found"})
	})
}
