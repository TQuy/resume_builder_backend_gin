package controllers

import (
	"fmt"
	"net/http"
	"resume_builder/go-gin-gorm/constants"
	"resume_builder/go-gin-gorm/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUser handle account registration
func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// Insert new user into database
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	models.DB.Create(&user)

	message := fmt.Sprintf("User %v created successfully", user.Username)
	c.JSON(http.StatusCreated, gin.H{"data": message})
}

// AuthUser handle account login
func AuthUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect username!"})
		return
	}
	// Compare password in database with password from request
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password!"})
	}
	// Generate token for request authentication
	jwtToken, err := generateJWT(user.ID, user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": jwtToken})
}

// generateJWT generate new JWT token
func generateJWT(id uint, username string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"nbf":      time.Now().Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(constants.SECRET_KEY))

	return tokenString, err
}
