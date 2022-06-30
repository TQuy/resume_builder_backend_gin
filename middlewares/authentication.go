package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// LoginRequired validate jwt token before passing it to handler
func LoginRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		var authorization []string = c.Request.Header["Authorization"]
		if len(authorization) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing token!"})
		}

		bearer := strings.Fields(authorization[0])
		if len(bearer) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect token format!"})
		}

		c.Keys = make(map[string]interface{})
		c.Keys["token"] = bearer[1]
		c.Next()
	}
}
