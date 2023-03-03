package main

import (
	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env.InitEnv()
	router := gin.Default()
	router.Static("/static", "./public")
	router.POST("/signup", handlers.SignUpHandler)

	router.Run(":8080")
}
