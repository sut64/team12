package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /InsuranceConverage
func CreateInsuranceConverage(c *gin.Context) {

	var insuranceconverages entity.InsuranceConverage
	var packageinsurs entity.PackageInsur
	var protections entity.Protection
	var totallists entity.Totallist
	var employees entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร insuranceconverage
	if err := c.ShouldBindJSON(&insuranceconverages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา package1 ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverages.PackageInsurID).First(&packageinsurs); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "packageinsur not found"})
		return
	}

	// 10: ค้นหา protection ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverages.ProtectionID).First(&protections); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "protection not found"})
		return
	}

	// 11: ค้นหา totallist ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverages.TotallistID).First(&totallists); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "totallist not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverages.EmployeeID).First(&employees); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 15: สร้าง Insuranceconverage
	ic := entity.InsuranceConverage{
		PackageInsur:   packageinsurs,                      // โยงความสัมพันธ์กับ Entity Package
		Protection:     protections,                        // โยงความสัมพันธ์กับ Entity Protection
		Totallist:      totallists,                         // โยงความสัมพันธ์กับ Entity Totallist
		Premium:        insuranceconverages.Premium,        // ตั้งค่าฟิลด์ Premium
		Protectiontype: insuranceconverages.Protectiontype, // ตั้งค่าฟิลด์ protectiontype
		Datetime:       insuranceconverages.Datetime,       // ตั้งค่าฟิลด์ Datetime
		Employee:       employees,
	}

	if _, err := govalidator.ValidateStruct(ic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 16: บันทึก
	if err := entity.DB().Create(&ic).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ic})
}

// GET /InsuranceConverage/:id
func GetInsuranceConverage(c *gin.Context) {
	var insuranceconverages entity.InsuranceConverage
	id := c.Param("id")
	if err := entity.DB().Preload("Totallist").Preload("Protection").Preload("PackageInsur").Preload("Employee").Raw("SELECT * FROM insurance_converages WHERE id = ?", id).Find(&insuranceconverages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": insuranceconverages})
}

// GET /insuranceconverages
func ListInsuranceConverages(c *gin.Context) {
	var insuranceconverages []entity.InsuranceConverage
	if err := entity.DB().Preload("Totallist").Preload("Protection").Preload("PackageInsur").Preload("Employee").Raw("SELECT * FROM insurance_converages").Find(&insuranceconverages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceconverages})
}

// DELETE /insuranceconverages/:id
func DeleteInsuranceConverage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM insurance_converages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insuranceconverage not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /insuranceconverages
func UpdateInsuranceConverage(c *gin.Context) {
	var insuranceconverages entity.InsuranceConverage
	if err := c.ShouldBindJSON(&insuranceconverages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", insuranceconverages.ID).First(&insuranceconverages); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insuranceconverage not found"})
		return
	}

	if err := entity.DB().Save(&insuranceconverages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceconverages})
}
