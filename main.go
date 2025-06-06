package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"stock-app/controller"
	"stock-app/db"
)

func main() {
	godotenv.Load()
	db.InitDB()

	router := gin.Default()
	controller.RegisterRoutes(router)
	router.Run(":8080")
}
