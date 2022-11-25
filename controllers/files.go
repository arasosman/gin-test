package controllers

import (
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindFiles(c *gin.Context) {
	var files []models.File
	models.DB.Find(&files)

	c.JSON(http.StatusOK, gin.H{"data": files})
}

func FindFile(c *gin.Context) { // Get model if exist
	var file models.File

	if err := models.DB.Where("id = ?", c.Param("id")).First(&file).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": file})
}

func CreateFile(c *gin.Context) {
	// Validate input
	var input CreateFileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var file models.File
	err := c.BindJSON(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Create(&file)

	c.JSON(http.StatusOK, gin.H{"data": file})
}

type CreateFileInput struct {
	Name           string `json:"name" binding:"required"`
	Path           string `json:"path"`
	FolderId       uint   `json:"folder_id" binding:"required"`
	InstallationId uint   `json:"installation_id" binding:"required"`
}
