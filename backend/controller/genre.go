package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

// GET /genre
// List all genre
func ListGenre(c *gin.Context) {
	var genre []entity.Genre
	if err := entity.DB().Raw("SELECT * FROM genre").Scan(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// GET /genre/:id
// Get genre by id
func GetGenre(c *gin.Context) {
	var genre entity.Genre
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM genre WHERE id = ?", id).Scan(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// POST /genre
func CreateGenre(c *gin.Context) {
	var genre entity.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// PATCH /genre
func UpdateGenre(c *gin.Context) {
	var genre entity.Genre
	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", genre.ID).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	if err := entity.DB().Save(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

// DELETE /genre/:id
func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM genre WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}
