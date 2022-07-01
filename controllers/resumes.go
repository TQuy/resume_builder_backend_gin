package controllers

import (
	"net/http"
	"resume_builder/go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

type CreateResumeInput struct {
	Name    string `json:"name" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// ListResumes list all the current user's resumes
func ListResumes(c *gin.Context) {
	var resumes []models.Resume
	models.DB.Where("author_id = ?", c.Keys["id"]).Select("id", "name", "content").Find(&resumes)
	c.JSON(http.StatusOK, gin.H{"data": resumes})
}

// CreateResume insert new record into resumes table
func CreateResume(c *gin.Context) {
	userID, _ := c.Keys["id"].(uint)

	var input CreateResumeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resumeExists bool
	if err := models.DB.Model(&models.Resume{}).Select("count(*)").Where("name = ? and author_id = ?", input.Name, userID).Find(&resumeExists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if resumeExists {
		c.JSON(http.StatusConflict, gin.H{"error": "Resume existed!"})
		return
	}

	resume := models.Resume{
		Name:     input.Name,
		Content:  input.Content,
		AuthorId: uint(userID),
	}
	models.DB.Create(&resume)
	c.JSON(http.StatusOK, gin.H{"data": resume})
}

func DelteResume(c *gin.Context) {
	var resume models.Resume
	if err := models.DB.Model(&models.Resume{}).Where("id = ?", c.Param("id")).First(&resume).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&resume)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
