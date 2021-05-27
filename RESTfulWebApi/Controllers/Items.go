package controllers

import (
	"net/http"

	"RestApi/models"

	"github.com/gin-gonic/gin"
)

type CreateItemInput struct {
	Name    string `json:"name" binding:"required"`
	Details string `json:"details" binding:"required"`
}

type UpdateItemInput struct {
	Name    string `json:"name"`
	Details string `json:"details"`
}

// GET/Api/Items
func GetItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

// GET/Api/Item/:id
func GetItemById(c *gin.Context) {
	var item models.Item

	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// POST/Api/Item/:id
func PostItem(c *gin.Context) {

	var input CreateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	item := models.Item{Name: input.Name, Details: input.Details}
	models.DB.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// PUT/Api/Item/:id
func PutItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": "Record not found!"})
		return
	}

	var input UpdateItemInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ERROR": err.Error()})
		return
	}

	models.DB.Model(&item).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": item})
}

// PUT/Api/Item/:id
func DeleteItem(c *gin.Context) {
	var item models.Item
	if err := models.DB.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&item)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
