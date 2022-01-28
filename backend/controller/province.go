package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

// GET /province
// List all provinces
func ListProvince(c *gin.Context) {
	var province []entity.Province
	if err := entity.DB().Raw("SELECT * FROM province").Scan(&province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

// GET /province/:id
// Get province by id
func GetProvince(c *gin.Context) {
	var province entity.Province
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM province WHERE id = ?", id).Scan(&province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

// POST /province
func CreateProvince(c *gin.Context) {
	var province entity.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

// PATCH /provinces
func UpdateProvince(c *gin.Context) {
	var province entity.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", province.ID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	if err := entity.DB().Save(&province).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

// DELETE /province/:id
func DeleteProvince(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM province WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
