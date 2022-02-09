package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/entity"
)

// POST /invoicepayment
func CreateInvoicePayment(c *gin.Context) {

	var invoicepayment entity.InvoicePayment
	var invoice entity.Invoice
	var customer entity.Customer
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร invoicepayment
	if err := c.ShouldBindJSON(&invoicepayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา invoice ด้วย id
	if tx := entity.DB().Where("id = ?", invoicepayment.InvoiceID).First(&invoice); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "locker not found"})
		return
	}

	// 10: ค้นหา customer ด้วย id
	if tx := entity.DB().Where("id = ?", invoicepayment.CustomerID).First(&customer); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "member not found"})
		return
	}

	// 11: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", invoicepayment.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}
	// 12: สร้าง InvoicePayment
	ip := entity.InvoicePayment{
		Invoice:       invoice,                      // โยงความสัมพันธ์กับ Entity Invoice
		Customer:      customer,                     // โยงความสัมพันธ์กับ Entity Customer
		Employee:      employee,                     // โยงความสัมพันธ์กับ Entity Employee
		PaymentTime:   invoicepayment.PaymentTime,   // ตั้งค่าฟิลด์ paymentTime
		InvoiceNumber: invoicepayment.InvoiceNumber, // ตั้งค่าฟิลด์ invoiceNumber
		PaymentAmount: invoicepayment.PaymentAmount, // ตั้งค่าฟิลด์ paymentAmount
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

// GET /invoicepayment/:id
func GetInvoicePayment(c *gin.Context) {
	var invoicepayment entity.InvoicePayment
	id := c.Param("id")
	if err := entity.DB().Preload("Invoice").Preload("Customer").Preload("Employee").Raw("SELECT * FROM invoice_payments WHERE id = ?", id).Find(&invoicepayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": invoicepayment})
}

// GET /invoice_payments
func ListInvoicePayments(c *gin.Context) {
	var invoicepayments []entity.InvoicePayment
	if err := entity.DB().Preload("Invoice").Preload("Customer").Preload("Employee").Raw("SELECT * FROM invoice_payments").Find(&invoicepayments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoicepayments})
}

// DELETE /invoice_payments/:id
func DeleteInvoicePayment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM invoice_payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoicepayment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /invoice_payments
func UpdateInvoicePayment(c *gin.Context) {
	var invoicepayment entity.InvoicePayment
	if err := c.ShouldBindJSON(&invoicepayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", invoicepayment.ID).First(&invoicepayment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoicepayment not found"})
		return
	}

	if err := entity.DB().Save(&invoicepayment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoicepayment})
}
