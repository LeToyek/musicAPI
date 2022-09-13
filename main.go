package main

import (
	"musicAPI/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/login", controllers.Login())

	router.Run("localhost:5000")
}
