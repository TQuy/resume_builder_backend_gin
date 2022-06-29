package main

import (
	"resume_builder/go-gin-gorm/controllers"
	"resume_builder/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/auth/register", controllers.CreateUser)

	r.Run("127.0.0.1:8080")
}
