package controller

import (
	"net/http"

	"github.com/Kaweethorn/team12/entity"
	"github.com/gin-gonic/gin"
)

// POST /invoices
func CreateInvoice(c *gin.Context) {
	var invoice entity.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

// GET /invoice/:id
func GetInvoice(c *gin.Context) {
	var invoice entity.Invoice
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM invoices WHERE id = ?", id).Scan(&invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}

// GET /invoices
func ListInvoices(c *gin.Context) {
	var invoices []entity.Invoice
	if err := entity.DB().Raw("SELECT * FROM invoices").Scan(&invoices).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoices})
}

// DELETE /invoices/:id
func DeleteInvoice(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM invoices WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /invoices
func UpdateInvoice(c *gin.Context) {
	var invoice entity.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", invoice.ID).First(&invoice); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invoice not found"})
		return
	}

	if err := entity.DB().Save(&invoice).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": invoice})
}
