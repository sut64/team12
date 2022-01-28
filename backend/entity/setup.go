package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("projectSE.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Employee{},
		&Hospitalnet{},
		&Genre{},
		&Status{},
		&Province{},
		&Invoice{},
		&Customer{},
		&InvoicePayment{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	// -- Employee Data
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายเอ",
		Email:    "aaa@example.com",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายบี",
		Email:    "bbb@example.com",
		Password: string(password),
	})

	var a Employee
	var b Employee
	db.Raw("Select * FROM employees WHERE email = ?", "aaa@example.com").Scan(&a)
	db.Raw("Select * FROM employees WHERE email = ?", "bbb@example.com").Scan(&b)

	// -- Invoice Data
	i1 := Invoice{
		PaymentType: "Credit",
	}
	db.Model(&Invoice{}).Create(&i1)

	i2 := Invoice{
		PaymentType: "Cash",
	}
	db.Model(&Invoice{}).Create(&i2)

	// -- Customer Data
	c := Customer{
		Name:         "นายซี",
		IdNumber:     "1329911111111",
		PolicyNumber: "00000001",
	}
	db.Model(&Customer{}).Create(&c)

	d := Customer{
		Name:         "นายดี",
		IdNumber:     "1329922222222",
		PolicyNumber: "00000002",
	}
	db.Model(&Customer{}).Create(&d)

	f := Customer{
		Name:         "นายเอฟ",
		IdNumber:     "1329933333333",
		PolicyNumber: "00000003",
	}
	db.Model(&Customer{}).Create(&f)
}
