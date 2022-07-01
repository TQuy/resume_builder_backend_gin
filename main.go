package main

import (
	"resume_builder/go-gin-gorm/controllers"
	"resume_builder/go-gin-gorm/middlewares"
	"resume_builder/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.POST("/auth/register", controllers.CreateUser)
	r.POST("/auth/login", controllers.AuthUser)

	r.GET("/resumes", middlewares.LoginRequired(), controllers.ListResumes)
	r.POST("/resumes", middlewares.LoginRequired(), controllers.CreateOrUpdateResume)

	r.GET("/resumes/:id", middlewares.LoginRequired(), controllers.FindResume)
	r.DELETE("/resumes/:id", middlewares.LoginRequired(), controllers.DeleteResume)

	r.Run("127.0.0.1:8000")
}
