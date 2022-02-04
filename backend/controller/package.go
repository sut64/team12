package controller

import (
	"net/http"

	"github.com/sut64/team12/entity"
	"github.com/gin-gonic/gin"
)

// POST /packages
func CreatePackage(c *gin.Context) {
	var package1 entity.Package
	if err := c.ShouldBindJSON(&package1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&package1).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": package1})
}

// GET /package/:id
func GetPackage(c *gin.Context) {
	var package1 entity.Package
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM packages WHERE id = ?", id).Find(&package1).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": package1})
}

// GET /package/watched/user/:id
func GetPackageWatchedByUser(c *gin.Context) {
	var package1 entity.Package
	id := c.Param("id")
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM packages WHERE owner_id = ? AND title = ?", id, "Watched").Find(&package1).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": package1})
}

// GET /packages
func ListPackages(c *gin.Context) {
	var packages []entity.Package
	if err := entity.DB().Preload("Owner").Raw("SELECT * FROM packages").Find(&packages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": packages})
}

// DELETE /packages/:id
func DeletePackage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM packages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "package not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /packages
func UpdatePackage(c *gin.Context) {
	var package1 entity.Package
	if err := c.ShouldBindJSON(&package1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", package1.ID).First(&package1); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "package not found"})
		return
	}

	if err := entity.DB().Save(&package1).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": package1})
}
