package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

func CreateInsuranceClaim(c *gin.Context) {
	var insuranceclaims entity.InsuranceClaim
	var motives entity.Motive
	var employee entity.Employee
	var customer entity.Customer

	if err := c.ShouldBindJSON(&insuranceclaims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceclaims.MotiveID).First(&motives); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "motive not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceclaims.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	if tx := entity.DB().Where("id = ?", insuranceclaims.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	ic := entity.InsuranceClaim{
		Patient:      insuranceclaims.Patient,
		Employee:     employee,
		Customer:     customer,
		Compensation: insuranceclaims.Compensation,
		Insdate:      insuranceclaims.Insdate,
		Motive:       motives,
	}
	if _, err := govalidator.ValidateStruct(ic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&ic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ic})
}

func GetInsuranceClaim(c *gin.Context) {
	var insuranceclaims entity.InsuranceClaim
	id := c.Param("id")
	if err := entity.DB().Preload("Employee").Preload("Motive").Preload("Customer").Raw("SELECT * FROM insurance_claims WHERE id = ?", id).Find(&insuranceclaims).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": insuranceclaims})
}

// GET /insuranceconverages
func ListInsuranceClaims(c *gin.Context) {
	var insuranceclaims []entity.InsuranceClaim
	if err := entity.DB().Preload("Employee").Preload("Motive").Preload("Customer").Raw("SELECT * FROM insurance_claims").Find(&insuranceclaims).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceclaims})
}

// DELETE /insuranceconverages/:id
func DeleteInsuranceClaim(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM insurance_claims WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insuranceclaims not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /invoice_payments
func UpdateInsuranceClaim(c *gin.Context) {
	var insuranceclaims entity.InsuranceClaim
	if err := c.ShouldBindJSON(&insuranceclaims); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceclaims.ID).First(&insuranceclaims); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insuranceconverage not found"})
		return
	}

	if err := entity.DB().Save(&insuranceclaims).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceclaims})
}
