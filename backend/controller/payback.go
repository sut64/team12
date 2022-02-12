package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /payback
func CreatePayback(c *gin.Context) {

	var paybacks entity.Payback
	var customers entity.Customer
	var employees entity.Employee
	var protections entity.Protection
	var banks entity.Bank

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร payback
	if err := c.ShouldBindJSON(&paybacks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 10: ค้นหา protection ด้วย id
	if tx := entity.DB().Where("id = ?", paybacks.ProtectionID).First(&protections); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "protection not found"})
		return
	}

	// 10: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", paybacks.CustomerID).First(&customers); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", paybacks.EmployeeID).First(&employees); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	// 11: ค้นหา bank ด้วย id
	if tx := entity.DB().Where("id = ?", paybacks.BankID).First(&banks); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bank not found"})
		return
	}
	// 12: สร้าง Payback
	ip := entity.Payback{
		Bank:       banks,
		Protection: protections,
		Customer:   customers, // โยงความสัมพันธ์กับ Entity Customer
		Employee:   employees, // โยงความสัมพันธ์กับ Entity Employee

		IDcard: paybacks.IDcard,
		Accout: paybacks.Accout,
		Year:   paybacks.Year,
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(ip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&ip).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ip})
}

// GET /payback/:id
func GetPayback(c *gin.Context) {
	var paybacks entity.Payback
	id := c.Param("id")
	if err := entity.DB().Preload("Bank").Preload("Customer").Preload("Employee").Preload("Protection").Raw("SELECT * FROM paybacks WHERE id = ?", id).Find(&paybacks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": paybacks})
}

// GET /payback
func ListPaybacks(c *gin.Context) {
	var paybacks []entity.Payback
	if err := entity.DB().Preload("Bank").Preload("Customer").Preload("Employee").Preload("Protection").Raw("SELECT * FROM paybacks").Find(&paybacks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paybacks})
}

// DELETE /paybacks/:id
func DeletePayback(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM paybacks WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payback not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /paybacks
func UpdatePayback(c *gin.Context) {
	var paybacks entity.Payback
	if err := c.ShouldBindJSON(&paybacks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", paybacks.ID).First(&paybacks); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payback not found"})
		return
	}

	if err := entity.DB().Save(&paybacks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paybacks})
}
