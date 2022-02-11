package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
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
	BuyInsurance    []Buyinsurance   `gorm:"foreignKey:CustomerID"`
	Paybacks        []Payback        `gorm:"foreignKey:CustomerID"`
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
type Bank struct {
	gorm.Model
	Name     string
	Bank     []Bank    `gorm:"foreignKey:GenreID"`
	Paybacks []Payback `gorm:"foreignKey:EmployeeID"`
}

type Employee struct {
	gorm.Model
	Name               string
	Email              string `gorm:"uniqueIndex"`
	Password           string
	Hospitalnet        []Hospitalnet        `gorm:"foreignKey:EmployeeID"`
	InvoicePayments    []InvoicePayment     `gorm:"foreignKey:EmployeeID"`
	BuyInsurance       []Buyinsurance       `gorm:"foreignKey:EmployeeID"`
	Paybacks           []Payback            `gorm:"foreignKey:EmployeeID"`
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:EmployeeID"`
}
type InvoicePayment struct {
	gorm.Model
	PaymentTime   time.Time `valid:"notpast~PaymentTime must not be in the past"`
	InvoiceNumber string    `valid:"matches(^[i]\\d{3}$)~InvoiceNumber is not correct"`
	PaymentAmount int       `valid:"IsPositive~PaymentAmount must be positive"`

	// InvoiceID ทำหน้าที่เป็น FK
	InvoiceID *uint
	Invoice   Invoice `gorm:"references:id" valid:"-"`

	// CustomerID ทำหน้าที่เป็น FK
	CustomerID *uint
	Customer   Customer `gorm:"references:id" valid:"-"`
	// EmployeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee `gorm:"references:id" valid:"-"`
}
type Hospitalnet struct {
	gorm.Model
	Name     string
	Contract float64 `valid:"IsPositive~Contract cannot be negative or 0"`
	Address  string `valid:"minstringlength(5)~Adddress should more than 5 charactor"`
	Adddate  time.Time `valid:"notpast~Date cannot be past"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	StatusID *uint
	Status   Status `gorm:"references:id"`

	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	GenreID *uint
	Genre   Genre `gorm:"references:id"`
}

type Buyinsurance struct {
	gorm.Model
	Consent      bool
	HealthInfrom string
	Adddate      time.Time

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`
}

type Payback struct {
	gorm.Model
	Name    string
	Year    float64
	Accout  int
	Address string `valid:"required"`

	CustomerID *uint
	Customer   Customer `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	BankID *uint
	Bank   Bank `gorm:"references:id"`
}

type Protection struct {
	gorm.Model
	Protection_name    string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:ProtectionID"`
}
type Package struct {
	gorm.Model
	Package_name       string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:PackageID"`
}
type Totallist struct {
	gorm.Model
	Totallist_cost     string
	InsuranceConverage []InsuranceConverage `gorm:"foreignKey:TotallistID"`
}

type InsuranceConverage struct {
	gorm.Model
	Protectiontype string
	Premium        uint
	Datetime       time.Time

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`

	ProtectionID *uint
	Protection   Protection `gorm:"references:id"`

	PackageID *uint
	Package   Package `gorm:"references:id"`

	TotallistID *uint
	Totallist   Totallist `gorm:"references:id"`
}

func init() {
    govalidator.CustomTypeTagMap.Set("notpast", func(i interface{} , context interface{}) bool{
        t := i.(time.Time)
		now := time.Now()
        return now.Before(t)
    })
}
	govalidator.CustomTypeTagMap.Set("IsPositive", func(i interface{}, context interface{}) bool {
		value := i.(int)
		return value >= 0
	})
}
