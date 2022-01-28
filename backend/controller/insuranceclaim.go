package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

func CreateInsuranceClaim(c *gin.Context) {
	var insuranceClaims entity.InsuranceClaim
	var motives entity.Motive
	var employee entity.Employee

	if err := c.ShouldBindJSON(&insuranceClaims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceClaims.MotiveID).First(&motives); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "motive not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceClaims.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	ic := entity.InsuranceClaim{
		Patient:      insuranceClaims.Patient,
		Employee:     employee,
		Compensation: insuranceClaims.Compensation,
		InsurDate:    insuranceClaims.InsurDate,
		Motive:       motives,
	}

	if err := entity.DB().Create(&ic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ic})
}

func GetAllInsuranceClaim(c *gin.Context) {
	var insuranceClaim []entity.InsuranceClaim
	if err := entity.DB().
		Preload("Motive").
		Raw("SELECT * FROM insuranceClaims").Find(&insuranceClaim).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceClaim})
}
