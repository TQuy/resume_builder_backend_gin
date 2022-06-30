package controllers

import (
	"net/http"
	"resume_builder/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

// ListResumes list all the current user's resumes
func ListResumes(c *gin.Context) {
	var resumes []models.Resume
	models.DB.Where("author_id = ?", c.Keys["id"]).Find(&resumes)
	c.JSON(http.StatusOK, gin.H{"data": resumes})
}
