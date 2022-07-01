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

// FindResume return specific resume based on resume's ID
func FindResume(c *gin.Context) {
	var resume models.Resume

	if err := models.DB.First(&resume, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": resume})
}

// CreateResume insert new record into resumes table
func CreateOrUpdateResume(c *gin.Context) {
	userID, _ := c.Keys["id"].(uint)

	var input CreateResumeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resume models.Resume
	result := models.DB.Where("name = ? and author_id = ?", input.Name, userID).Limit(1).Find(&resume)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	// If resume found, update it
	if result.RowsAffected == 1 {
		UpdateResume(c, input, resume)
		return
	}
	// else create new
	resume = models.Resume{
		Name:     input.Name,
		Content:  input.Content,
		AuthorId: uint(userID),
	}
	models.DB.Create(&resume)
	c.JSON(http.StatusOK, gin.H{"data": resume})
}

// UpdateResume handle resume update
func UpdateResume(c *gin.Context, input CreateResumeInput, resume models.Resume) {
	resume.Content = input.Content
	models.DB.Save(&resume)
	c.JSON(http.StatusOK, gin.H{"data": resume})
}

// Delete specific resume based on resume's ID
func DeleteResume(c *gin.Context) {
	var resume models.Resume
	if err := models.DB.Model(&models.Resume{}).Where("id = ?", c.Param("id")).First(&resume).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&resume)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
