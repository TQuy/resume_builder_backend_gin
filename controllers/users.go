package controllers

import (
	"fmt"
	"net/http"
	"resume_builder/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}
	models.DB.Create(&user)

	message := fmt.Sprintf("User %v created successfully", user.Username)

	c.JSON(http.StatusOK, gin.H{"data": message})
}
