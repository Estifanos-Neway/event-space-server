package main

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	env.InitEnv()
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/static", "./static")
	router.POST("/signup", handlers.SignUpHandler)
	router.POST("/verify-signup", handlers.VerifySignUpHandler)
	router.POST("/signin", handlers.SignInHandler)
	router.POST("/refresh", handlers.RefreshHandler)
	router.POST("/signout", handlers.SignOutHandler)
	router.POST("/ticket-sell", handlers.TicketSellHandler)
	router.POST("/file-upload", handlers.FileUploadHandler)

	log.Println("Starting...")
	router.Run(":8080")
}
