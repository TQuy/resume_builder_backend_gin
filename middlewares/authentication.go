package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"resume_builder/go-gin-gorm/constants"
	"resume_builder/go-gin-gorm/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// LoginRequired validate and parse jwt token before passing it to handler
func LoginRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		var authorization []string = c.Request.Header["Authorization"]
		if len(authorization) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token!"})
			c.Abort()
			return
		}
		// validate and parse jwt token
		userID, err := getAccountID(authorization)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			fmt.Println("-----------stop here")
			c.Abort()
			return
		}
		// Save userID to request session
		c.Keys = make(map[string]interface{})
		fmt.Printf("------------------------------userID: %v\n", userID)
		c.Keys["id"] = userID
		// c.Next()
	}
}

func getAccountID(authorization []string) (uint, error) {
	bearer := strings.Fields(authorization[0])
	if len(bearer) < 2 {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect token format!"})
		return 0, errors.New("Incorrect token format!")
	}

	tokenString := bearer[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(constants.SECRET_KEY), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Unexpected error!")
	}
	// This shit need type assertion
	// its type is float64, but convert it to string first before convert to uint64 to make sure
	userID, err := strconv.ParseUint(fmt.Sprint(claims["id"]), 10, 0)
	if err != nil {
		return 0, err
	}

	var userExists bool
	if err = models.DB.Model(&models.User{}).
		Select("count(*) > 0").
		Where("id = ?", userID).
		Find(&userExists).
		Error; err != nil {
		return 0, err
	}

	if !userExists {
		return 0, errors.New("User not found!")
	}

	return uint(userID), nil
}
