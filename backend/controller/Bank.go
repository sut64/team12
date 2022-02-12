package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /banks
func CreateBank(c *gin.Context) {
	var bank entity.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&bank).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bank})
}

// GET /bank/:id
func GetBank(c *gin.Context) {
	var bank entity.Bank
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM banks WHERE id = ?", id).Scan(&bank).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bank})
}

// GET /banks
func ListBanks(c *gin.Context) {
	var banks []entity.Bank
	if err := entity.DB().Raw("SELECT * FROM banks").Scan(&banks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": banks})
}

// DELETE /banks/:id
func DeleteBank(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM banks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /banks
func UpdateBank(c *gin.Context) {
	var bank entity.Bank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", bank.ID).First(&bank); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
		return
	}

	if err := entity.DB().Save(&bank).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bank})
}
