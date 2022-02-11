package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

// POST /protections
func CreateProtection(c *gin.Context) {
	var protection entity.Protection
	if err := c.ShouldBindJSON(&protection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&protection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": protection})
}

// GET /protection/:id
func GetProtection(c *gin.Context) {
	var protection entity.Protection
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM protections WHERE id = ?", id).Scan(&protection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": protection})
}

// GET /protections
func ListProtections(c *gin.Context) {
	var protections []entity.Protection
	if err := entity.DB().Raw("SELECT * FROM protections").Scan(&protections).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": protections})
}

// DELETE /protections/:id
func DeleteProtection(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM totallists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "protection not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /protections
func UpdateProtection(c *gin.Context) {
	var protection entity.Protection
	if err := c.ShouldBindJSON(&protection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", protection.ID).First(&protection); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "protection not found"})
		return
	}

	if err := entity.DB().Save(&protection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": protection})
}
