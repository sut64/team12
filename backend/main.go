package main

import (
	"github.com/B5808724/team12/controller"
	"github.com/B5808724/team12/entity"
	"github.com/B5808724/team12/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Protection Routes
			protected.GET("/protections", controller.ListProtections)
			protected.GET("/protection/:id", controller.GetProtection)
			protected.POST("/protections", controller.CreateProtection)
			protected.PATCH("/protections", controller.UpdateProtection)
			protected.DELETE("/protections/:id", controller.DeleteProtection)

			// Package Routes
			protected.GET("/packages", controller.ListPackages)
			protected.GET("/package/:id", controller.GetPackage)
			protected.POST("/packages", controller.CreatePackage)
			protected.PATCH("/packages", controller.UpdatePackage)
			protected.DELETE("/packages/:id", controller.DeletePackage)

			// Totallist Routes
			protected.GET("/totallists", controller.ListTotallists)
			protected.GET("/totallist/:id", controller.GetTotallist)
			protected.PATCH("/totallists", controller.UpdateTotallist)
			protected.DELETE("/totallists/:id", controller.DeleteTotallist)

			
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
