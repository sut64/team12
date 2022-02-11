package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /totallists
func CreateTotallist(c *gin.Context) {
	var totallist entity.Totallist
	if err := c.ShouldBindJSON(&totallist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&totallist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": totallist})
}

// GET /totallist/:id
func GetTotallist(c *gin.Context) {
	var totallist entity.Totallist
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM totallists WHERE id = ?", id).Scan(&totallist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": totallist})
}

// GET /totallists
func ListTotallists(c *gin.Context) {
	var totallists []entity.Totallist
	if err := entity.DB().Raw("SELECT * FROM totallists").Scan(&totallists).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": totallists})
}

// DELETE /totallists/:id
func DeleteTotallist(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM totallists WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "totallist not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /totallists
func UpdateTotallist(c *gin.Context) {
	var totallist entity.Totallist
	if err := c.ShouldBindJSON(&totallist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", totallist.ID).First(&totallist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "totallist not found"})
		return
	}

	if err := entity.DB().Save(&totallist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": totallist})
}
