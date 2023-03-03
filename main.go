package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Static("/static", "./public")

	router.Run(":8080")
}
