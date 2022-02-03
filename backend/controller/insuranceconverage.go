package controller

import (
	"net/http"

	"github.com/sut64/team12/entity"
	"github.com/gin-gonic/gin"
)

// POST /InsuranceConverage
func CreateInsuranceConverage(c *gin.Context) {

	var insuranceconverage entity.InsuranceConverage
	var package1 entity.Package
	var protection entity.Protection
	var totallist entity.Totallist
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร insuranceconverage
	if err := c.ShouldBindJSON(&insuranceconverage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา package1 ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverage.PackageID).First(&package1); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "package not found"})
		return
	}

	// 10: ค้นหา protection ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverage.ProtectionID).First(&protection); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "protection not found"})
		return
	}

	// 11: ค้นหา totallist ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverage.TotallistID).First(&totallist); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "totallist not found"})
		return
	}
	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", insuranceconverage.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 15: สร้าง Insuranceconverage
	ic := entity.InsuranceConverage{
		Package:       package1,                      // โยงความสัมพันธ์กับ Entity Package
		Protection:      protection,                     // โยงความสัมพันธ์กับ Entity Protection
		Totallist:      totallist,                     // โยงความสัมพันธ์กับ Entity Totallist
		Employee:      employee,                     // โยงความสัมพันธ์กับ Entity Employee
		Premium:   insuranceconverage.Premium,   // ตั้งค่าฟิลด์ Premium
		Protectiontype : insuranceconverage.Protectiontype, // ตั้งค่าฟิลด์ protectiontype
		Datetime: insuranceconverage.Datetime, // ตั้งค่าฟิลด์ Datetime
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
	var insuranceconverage entity.InsuranceConverage
	id := c.Param("id")
	if err := entity.DB().Preload("Totallist").Preload("Protection").Preload("Package").Preload("Employee").Raw("SELECT * FROM insuranceconverages WHERE id = ?", id).Find(&insuranceconverage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": insuranceconverage})
}

// GET /insuranceconverages
func ListInsuranceConverages(c *gin.Context) {
	var insuranceconverages []entity.InsuranceConverage
	if err := entity.DB().Preload("Totallist").Preload("Protection").Preload("Package").Preload("Employee").Raw("SELECT * FROM invoice_payments").Find(&insuranceconverages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": insuranceconverages})
}

// DELETE /insuranceconverages/:id
func DeleteInsuranceConverage(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM insuranceconverages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insuranceconverage not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /invoice_payments
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
