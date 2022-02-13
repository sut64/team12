package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /watch_videos
func CreateBuyinsurance(c *gin.Context) {

	var Buyinsurance entity.Buyinsurance
	var Customer entity.Customer
	var Employee entity.Employee
	var insuranceconverages entity.InsuranceConverage

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร buyinsurance
	if err := c.ShouldBindJSON(&Buyinsurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา insuranceconverage ด้วย id
	if tx := entity.DB().Where("id = ?", Buyinsurance.InsuranceConverageID).First(&insuranceconverages); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insurance not found"})
		return
	}

	// 10: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", Buyinsurance.CustomerID).First(&Customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", Buyinsurance.EmployeeID).First(&Employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found"})
		return
	}
	// 12: สร้าง Buyinsurance
	bi := entity.Buyinsurance{
		Customer:           Customer,            // โยงความสัมพันธ์กับ Entity Customer
		Employee:           Employee,            // โยงความสัมพันธ์กับ Entity Employee
		InsuranceConverage: insuranceconverages, // โยงความสัมพันธ์กับ Entity InsuranceConverage
		Adddate:            Buyinsurance.Adddate,
		Consent:            Buyinsurance.Consent,
		Healthinfrom:       Buyinsurance.Healthinfrom,
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(bi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&bi).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": bi})
}

// GET /buyinsuranceconverage/:id
func GetBuyinsurance(c *gin.Context) {
	var Buyinsurance entity.Buyinsurance
	id := c.Param("id")
	if err := entity.DB().Preload("InsuranceConverage").Preload("Customer").Preload("Employee").Raw("SELECT * FROM Buyinsurances WHERE id = ?", id).Find(&Buyinsurance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Buyinsurance})
}

// GET /Buyinsurances
func ListBuyinsurances(c *gin.Context) {
	var Buyinsurances []entity.Buyinsurance
	if err := entity.DB().Preload("InsuranceConverage").Preload("Customer").Preload("Employee").Raw("SELECT * FROM Buyinsurances").Find(&Buyinsurances).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Buyinsurances})
}

// DELETE /watch_videos/:id
func DeleteBuyinsurance(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Buyinsurances WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "buyinsurance not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Buyinsurance
func UpdateBuyinsurance(c *gin.Context) {
	var Buyinsurance entity.Buyinsurance
	if err := c.ShouldBindJSON(&Buyinsurance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", Buyinsurance.ID).First(&Buyinsurance); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "buyinsurance not found"})
		return
	}

	if err := entity.DB().Save(&Buyinsurance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Buyinsurance})
}
