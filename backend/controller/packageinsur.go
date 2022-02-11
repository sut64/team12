package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /package
func CreatePackageInsur(c *gin.Context) {
	var packageinsur entity.PackageInsur
	if err := c.ShouldBindJSON(&packageinsur); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&packageinsur).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": packageinsur})
}

// GET /package/:id
func GetPackageInsur(c *gin.Context) {
	var packageinsur entity.PackageInsur
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM package_insurs WHERE id = ?", id).Find(&packageinsur).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": packageinsur})
}

// GET /packages
func ListPackageInsurs(c *gin.Context) {
	var packageinsurs []entity.PackageInsur
	if err := entity.DB().Raw("SELECT * FROM package_insurs").Find(&packageinsurs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": packageinsurs})
}

// DELETE /packages/:id
func DeletePackageInsur(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM package_insurs WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "package not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /packages
func UpdatePackageInsur(c *gin.Context) {
	var packageinsur entity.PackageInsur
	if err := c.ShouldBindJSON(&packageinsur); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", packageinsur.ID).First(&packageinsur); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "package not found"})
		return
	}

	if err := entity.DB().Save(&packageinsur).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": packageinsur})
}
