package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sut64/team12/controller"
	"github.com/sut64/team12/entity"
	"github.com/sut64/team12/middlewares"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Invoice Routes
			protected.GET("/invoices", controller.ListInvoices)
			protected.GET("/invoice/:id", controller.GetInvoice)
			protected.POST("/invoices", controller.CreateInvoice)
			protected.PATCH("/invoices", controller.UpdateInvoice)
			protected.DELETE("/invoices/:id", controller.DeleteInvoice)

			// Customer Routes
			protected.GET("/customers", controller.ListCustomers)
			protected.GET("/customer/:id", controller.GetCustomer)
			protected.POST("/customers", controller.CreateCustomer)
			protected.PATCH("/customers", controller.UpdateCustomer)
			protected.DELETE("/customers/:id", controller.DeleteCustomer)

			// Employee Routes
			protected.GET("/employees", controller.ListEmployees)
			protected.GET("/employee/:id", controller.GetEmployee)
			protected.PATCH("/employees", controller.UpdateEmployee)
			protected.DELETE("/employees/:id", controller.DeleteEmployee)

			// InvoicePayment Routes
			protected.GET("/invoice_payments", controller.ListInvoicePayments)
			protected.GET("/invoicepayment/:id", controller.GetInvoicePayment)
			protected.POST("/invoice_payments", controller.CreateInvoicePayment)
			protected.PATCH("/invoice_payments", controller.UpdateInvoicePayment)
			protected.DELETE("/invoice_payments/:id", controller.DeleteInvoicePayment)

			//BuyInsurance Routes
			protected.GET("/buyinsurances", controller.ListBuyinsurances)
			protected.GET("/buyinsurance/:id", controller.GetBuyinsurance)
			protected.POST("/buyinsurances", controller.CreateBuyinsurance)
			protected.PATCH("/buyinsurances", controller.UpdateBuyinsurance)
			protected.DELETE("/buyinsurances/:id", controller.DeleteBuyinsurance)

			// Genre Routes
			protected.GET("/genre", controller.ListGenre)
			protected.GET("/genre/:id", controller.GetGenre)
			protected.POST("/genre", controller.CreateGenre)
			protected.PATCH("/genre", controller.UpdateGenre)
			protected.DELETE("/genre/:id", controller.DeleteGenre)	
			
			// Status Routes
			protected.GET("/status", controller.ListStatus)
			protected.GET("/status/:id", controller.GetStatus)
			protected.POST("/status", controller.CreateStatus)
			protected.PATCH("/status", controller.UpdateStatus)
			protected.DELETE("/status/:id", controller.DeleteStatus)
			// Province Routes
			protected.GET("/province", controller.ListProvince)
			protected.GET("/province/:id", controller.GetProvince)
			protected.POST("/province", controller.CreateProvince)
			protected.PATCH("/provinces", controller.UpdateProvince)
			protected.DELETE("/province/:id", controller.DeleteProvince)

			// Hospital Routes
			protected.GET("/hospitalnets", controller.ListHospitalnets)
			protected.GET("/hospitalnet/id", controller.GetHospitalnet)
			protected.POST("/hospitalnets", controller.CreateHospitalnet)
			protected.PATCH("/hospitalnet", controller.UpdateHospitalnets)
			protected.DELETE("/hospitalnet/id", controller.DeleteHospitalnets)
			
			protected.GET("/protections", controller.ListProtections)
			protected.GET("/protection/:id", controller.GetProtection)
			protected.POST("/protections", controller.CreateProtection)
			protected.PATCH("/protections", controller.UpdateProtection)
			protected.DELETE("/protections/:id", controller.DeleteProtection)

			// Package Routes
			protected.GET("/packageinsurs", controller.ListPackageInsurs)
			protected.GET("/packageinsur/:id", controller.GetPackageInsur)
			protected.POST("/packageinsurs", controller.CreatePackageInsur)
			protected.PATCH("/packageinsurs", controller.UpdatePackageInsur)
			protected.DELETE("/packageinsurs/:id", controller.DeletePackageInsur)

			// Totallist Routes
			protected.GET("/totallists", controller.ListTotallists)
			protected.GET("/totallist/:id", controller.GetTotallist)
			protected.POST("/totallists", controller.CreateTotallist)
			protected.PATCH("/totallists", controller.UpdateTotallist)
			protected.DELETE("/totallists/:id", controller.DeleteTotallist)

			// InsuranceConverage Routes
			protected.GET("/insuranceconverages", controller.ListInsuranceConverages)
			protected.GET("/insuranceconverage/:id", controller.GetInsuranceConverage)
			protected.POST("/insuranceconverages", controller.CreateInsuranceConverage)
			protected.PATCH("/insuranceconverages", controller.UpdateInsuranceConverage)
			protected.DELETE("/insuranceconverages/:id", controller.DeleteInsuranceConverage)

			protected.GET("/motives", controller.ListMotives)
			protected.GET("/motive:id", controller.GetMotive)
			protected.POST("/motive", controller.CreateMotive)
			protected.PATCH("/motives", controller.UpdateMotive)
			protected.DELETE("/motive/:id", controller.DeleteMotive)

			protected.GET("/insuranceclaims", controller.ListInsuranceClaims)
			protected.GET("/insuranceclaim/:id", controller.GetInsuranceClaim)
			protected.POST("/insuranceclaim", controller.CreateInsuranceClaim)
			protected.PATCH("/insuranceclaims", controller.UpdateInsuranceClaim)
			protected.DELETE("/insuranceclaim/:id", controller.DeleteInsuranceClaim)

		}
	}

	// Employee Routes
	r.POST("/employees", controller.CreateEmployee)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
