package main

import (
	"rndev/backend-api/config"
	"rndev/backend-api/database"
	"rndev/backend-api/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()
	//setup router
	//inisialiasai Gin
	r := routes.SetupRouter()
	
	
	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}

// func SetupRouter() *gin.Engine{
// 	router := gin.Default()

// 	router.POST("/api/register", controllers.Register)
// 	return router
// }