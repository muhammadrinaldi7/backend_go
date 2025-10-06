package routes

import (
	"rndev/backend-api/controllers"
	"rndev/backend-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// route register
	router.POST("/api/register", controllers.Register)
	router.POST("/api/login", controllers.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	return router
}