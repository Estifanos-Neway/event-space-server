package main

import (
	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Aaa int = 8

func main() {
	godotenv.Load()
	env.InitEnv()
	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/static", "./public")
	router.POST("/signup", handlers.SignUpHandler)
	router.POST("/verify-signup", handlers.VerifySignUpHandler)
	router.POST("/signin", handlers.SignInHandler)
	router.POST("/refresh", handlers.RefreshHandler)

	router.Run(":8080")
}
