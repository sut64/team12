package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /hospitalnet
func CreateHospitalnet(c *gin.Context) {

	var hospitalnet entity.Hospitalnet
	var status entity.Status
	var province entity.Province
	var genre entity.Genre
	var employee entity.Employee

	if err := c.ShouldBindJSON(&hospitalnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnet.GenreID).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	if _, err := govalidator.ValidateStruct(hospitalnet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hn := entity.Hospitalnet{
		Genre:         	genre,                     
		Province:      	province,                    
		Employee:      	employee,                     
		Status:		   	status,
		Adddate:   		hospitalnet.Adddate,
		Name:			hospitalnet.Name,   			
		Address: 		hospitalnet.Address,
		Contract: 		hospitalnet.Contract,
	}
	
	if err := entity.DB().Create(&hn).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hn})
}


func GetHospitalnet(c *gin.Context) {
	var hospitalnet entity.Hospitalnet
	id := c.Param("id")
	if err := entity.DB().Preload("Genre").Preload("Status").Preload("Employee").Preload("Province").Raw("SELECT * FROM hospitalnets WHERE id = ?", id).Find(&hospitalnet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": hospitalnet})
}

// GET /hospitalnets
func ListHospitalnets(c *gin.Context) {
	var hospitalnets []entity.Hospitalnet
	if err := entity.DB().Preload("Genre").Preload("Status").Preload("Employee").Preload("Province").Raw("SELECT * FROM hospitalnets").Find(&hospitalnets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hospitalnets})
}

// DELETE /hospitalnets/:id
func DeleteHospitalnets(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM hospitalnets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hospitalnet not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /hospitalnetss
func UpdateHospitalnets(c *gin.Context) {
	var hospitalnets entity.Hospitalnet
	if err := c.ShouldBindJSON(&hospitalnets); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", hospitalnets.ID).First(&hospitalnets); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "hospitalnet not found"})
		return
	}

	if err := entity.DB().Save(&hospitalnets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": hospitalnets})
}
