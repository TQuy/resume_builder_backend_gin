package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ListResumes(c *gin.Context) {
	fmt.Println("--------------------------resume", c.Keys)
}
