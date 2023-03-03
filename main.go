package main

import (
	"fmt"

	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/handlers"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	// router.Run(":8080")

	type emailVerificationClaims struct {
		User types.User `json:"user"`
		jwt.RegisteredClaims
	}
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImVtYWlsIjoiZXN0aWZhbm9zLm5ld2F5LmRAZ21haWwuY29tIiwibmFtZSI6IlNvbWUgT25lIiwicGFzc3dvcmRIYXNoIjoiXHVmZmZkXHVmZmZkXHVmZmZkSFtmXHVmZmZkXHVmZmZkXHVmZmZkb3RcdWZmZmRcdWZmZmQ0XHVmZmZkXHVmZmZkXHVmZmZk4oaNIX9cdWZmZmRHMVVUIEpcdWZmZmRcdWZmZmRcdWZmZmQifX0.5sys4T36tjsl6r4PhsmbO5b23KTZKU5csLOPoYqzfvE"
	token, err := jwt.ParseWithClaims(tokenString, &emailVerificationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env.JWT_SECRETE), nil
	})
	if claims, ok := token.Claims.(*emailVerificationClaims); ok && token.Valid {
		fmt.Println(claims)
	} else {
		fmt.Println(err)
	}
	// log.Println("")
	// log.Println(token)
	// log.Println(err)
}
