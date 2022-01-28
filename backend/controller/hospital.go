package controller

import (
	"github.com/Kaweethorn/team12/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /staffs
func CreateHospital(c *gin.Context) {
	var hospitalnet entity.Hospitalnet
	var employee entity.Employee
	var status entity.Status
	var genre entity.Genre
	var province entity.Province

	if err := c.ShouldBindJSON(&hospitalnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.GenreID).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Genre not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Province not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	rif := entity.Hospitalnet{
		Employee: employee,
		Status:   status,
		Genre:    genre,
		Province: province,
		Name:     hospitalnet.Name,
		Contract: hospitalnet.Contract,
		Adddate:  hospitalnet.Adddate,
		Address:  hospitalnet.Address,
	}

	if err := entity.DB().Create(&rif).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": rif})
}

//GET ("/hospital/id")
//GET by id
func GetHospital(c *gin.Context) {

	var hospital entity.Hospitalnet

	if err := entity.DB().Preload("Employee").Preload("Status").Preload("Genre").Preload("Province").Raw("SELECT * FROM Hospital").Find(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hospital})
}

//GET ("/hospital")
func ListHospital(c *gin.Context) {

	var hospitals []entity.Hospitalnet

	if err := entity.DB().Preload("Employee").Preload("Status").Preload("Genre").Preload("Province").Find(&hospitals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hospitals})
}

// DELETE ("/hospital/id")
// DELETE by id
func DeleteHospital(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Hospital WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hospital not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH ("/hospital")
func UpdateHospital(c *gin.Context) {

	var hospital entity.Hospitalnet

	if err := c.ShouldBindJSON(&hospital); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", hospital.ID).First(&hospital); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hospital not found"})
		return
	}
	if err := entity.DB().Save(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hospital})
}
