package entity

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	PaymentType     string
	InvoicePayments []InvoicePayment `gorm:"foreignKey:InvoiceID"`
}
type Customer struct {
	gorm.Model
	Name            string
	IdNumber        string           `gorm:"uniqueIndex"`
	PolicyNumber    string           `gorm:"uniqueIndex"`
	InvoicePayments []InvoicePayment `gorm:"foreignKey:CustomerID"`
}
type Status struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:StatusID"`
}
type Province struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:ProvinceID"`
}
type Genre struct {
	gorm.Model
	Name        string
	Hospitalnet []Hospitalnet `gorm:"foreignKey:GenreID"`
}

type Employee struct {
	gorm.Model
	Name            string
	Email           string `gorm:"uniqueIndex"`
	Password        string
	Hospitalnet     []Hospitalnet    `gorm:"foreignKey:EmployeeID"`
	InvoicePayments []InvoicePayment `gorm:"foreignKey:CustomerID"`
}
type InvoicePayment struct {
	gorm.Model
	PaymentTime   time.Time
	InvoiceNumber string
	PaymentAmount int

	// InvoiceID ทำหน้าที่เป็น FK
	InvoiceID *uint
	Invoice   Invoice

	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer
	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee
}
type Hospitalnet struct {
	gorm.Model
	Name     string
	Contract float64
	Address  string
	Adddate  time.Time

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	GenreID *uint
	Genre   Genre `gorm:"references:id"`
}
