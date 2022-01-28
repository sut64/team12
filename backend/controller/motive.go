package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

// POST /resolutions
func CreateMotive(c *gin.Context) {
	var motive entity.Motive
	if err := c.ShouldBindJSON(&motive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&motive).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": motive})
}

// GET /resolution/:id
func GetMotive(c *gin.Context) {
	var motive entity.Motive
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM motives WHERE id = ?", id).Scan(&motive).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": motive})
}

// GET /resolutions
func ListMotives(c *gin.Context) {
	var motives []entity.Motive
	if err := entity.DB().Raw("SELECT * FROM motives").Scan(&motives).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": motives})
}

// DELETE /resolutions/:id
func DeleteMotive(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM motives WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "motive not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /resolutions
func UpdateMotive(c *gin.Context) {
	var motive entity.Motive
	if err := c.ShouldBindJSON(&motive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", motive.ID).First(&motive); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "motive not found"})
		return
	}

	if err := entity.DB().Save(&motive).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": motive})
}
